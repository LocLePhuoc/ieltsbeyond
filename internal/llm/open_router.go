package llm

import (
	"net/http"
)

var httpClient = &http.Client{}

func NewOpenRouterProvider(baseUrl string, apiKey string, model string, timeout int) *OpenAICompatProvider {
	return &OpenAICompatProvider{
		APIKey:     apiKey,
		BaseURL:    baseUrl,
		Model:      model,
		Timeout:    timeout,
		HTTPClient: &http.Client{},
	}
}
