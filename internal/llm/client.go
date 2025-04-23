package llm

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

// NewClient
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
