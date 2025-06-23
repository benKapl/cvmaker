package llm

import (
	"net/http"
	"time"
)

type baseClient struct {
	baseUrl    string
	httpClient http.Client
	apiKey     string
}

func newBaseClient(baseUrl, apiKey string, timeout time.Duration) baseClient {
	return baseClient{
		baseUrl: baseUrl,
		httpClient: http.Client{
			Timeout: timeout,
		},
		apiKey: apiKey,
	}
}
