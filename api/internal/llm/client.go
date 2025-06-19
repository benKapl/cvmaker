package llm

import (
	"net/http"
	"time"
)

type LLMClient struct {
	baseUrl    string
	httpClient http.Client
	llmConfig  LLMConfig
}

func NewClient(timeout time.Duration) LLMClient {
	return LLMClient{
		baseUrl: baseUrl,
		httpClient: http.Client{
			Timeout: timeout,
		},
		llmConfig: NewLLMConfig(),
	}
}
