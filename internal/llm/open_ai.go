package llm

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OpenAICompatProvider struct {
	APIKey       string
	Model        string
	BaseURL      string // "https://openrouter.ai/api/v1"
	HTTPClient   *http.Client
	ExtraHeaders map[string]string
	Timeout      int // ví dụ HTTP-Referer, X-Title cho OpenRouter
}

func (p *OpenAICompatProvider) Complete(ctx context.Context, systemPrompt string, parts []UserPart) (string, error) {
	type imageURL struct {
		URL string `json:"url"`
	}
	type contentPart struct {
		Type     string    `json:"type"`
		Text     string    `json:"text,omitempty"`
		ImageURL *imageURL `json:"image_url,omitempty"`
	}
	type msg struct {
		Role    string      `json:"role"`
		Content interface{} `json:"content"`
	}

	var userContent []contentPart
	for _, part := range parts {
		if part.IsImage() {
			b64 := base64.StdEncoding.EncodeToString(part.ImageData)
			uri := fmt.Sprintf("data:%s;base64,%s", part.ImageMIME, b64)
			userContent = append(userContent, contentPart{Type: "image_url", ImageURL: &imageURL{URL: uri}})
		} else {
			userContent = append(userContent, contentPart{Type: "text", Text: part.Text})
		}
	}

	reqBody := map[string]interface{}{
		"model":              p.Model,
		"max_tokens":         4000,
		"temperature":        0.2,
		"repetition_penalty": 1.2,
		"response_format":    map[string]string{"type": "json_object"},
		"messages": []msg{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userContent},
		},
	}

	payload, _ := json.Marshal(reqBody)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(p.Timeout)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.BaseURL, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+p.APIKey)
	req.Header.Set("Content-Type", "application/json")
	for k, v := range p.ExtraHeaders {
		req.Header.Set(k, v)
	}

	resp, err := p.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request failed: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var parsed struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("Response parse failed: %w (body: %s)", err, string(body))
	}
	if parsed.Error != nil {
		return "", fmt.Errorf("Response error: %s", parsed.Error.Message)
	}
	if len(parsed.Choices) == 0 {
		return "", fmt.Errorf("Empty response (HTTP %d): %s", resp.StatusCode, string(body))
	}
	return parsed.Choices[0].Message.Content, nil
}
