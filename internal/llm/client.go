package llm

import (
	"net/http"
	"time"
)

type Client struct {
	baseUrl    string
	httpClient http.Client
	LLMConfig  LLMConfig
}

func NewClient(timeout time.Duration) Client {
	return Client{
		baseUrl: baseUrl,
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
