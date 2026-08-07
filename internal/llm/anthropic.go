package llm

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type AnthropicCompatProvider struct {
	APIKey     string
	BaseURL    string
	Model      string
	HTTPClient *http.Client
	Timeout    int
}

func (p *AnthropicCompatProvider) Complete(ctx context.Context, systemPrompt string, parts []UserPart) (string, error) {
	type source struct {
		Type      string `json:"type"`
		MediaType string `json:"media_type,omitempty"`
		Data      string `json:"data,omitempty"`
	}
	type block struct {
		Type   string  `json:"type"`
		Text   string  `json:"text,omitempty"`
		Source *source `json:"source,omitempty"`
	}

	var content []block
	for _, part := range parts {
		if part.IsImage() {
			b64 := base64.StdEncoding.EncodeToString(part.ImageData)
			content = append(content, block{
				Type:   "image",
				Source: &source{Type: "base64", MediaType: part.ImageMIME, Data: b64},
			})
		} else {
			content = append(content, block{Type: "text", Text: part.Text})
		}
	}

	reqBody := map[string]interface{}{
		"model":       p.Model,
		"max_tokens":  2000,
		"temperature": 0.2,
		"system":      systemPrompt,
		"messages": []map[string]interface{}{
			{"role": "user", "content": content},
		},
	}

	payload, _ := json.Marshal(reqBody)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(p.Timeout)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.BaseURL, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("x-api-key", p.APIKey)
	// req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("content-type", "application/json")

	resp, err := p.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request failed: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var parsed struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", fmt.Errorf("Parse resposne fail: %w (body: %s)", err, string(body))
	}
	if parsed.Error != nil {
		return "", fmt.Errorf("Response error : %s", parsed.Error.Message)
	}

	var sb strings.Builder
	for _, c := range parsed.Content {
		if c.Type == "text" {
			sb.WriteString(c.Text)
		}
	}
	if sb.Len() == 0 {
		return "", fmt.Errorf("Empty response (HTTP %d): %s", resp.StatusCode, string(body))
	}
	return sb.String(), nil
}
