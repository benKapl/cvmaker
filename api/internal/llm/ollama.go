package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ollamaGenerateParams struct {
	Model    string         `json:"model"`
	Prompt   string         `json:"prompt"`
	Format   map[string]any `json:"format,omitempty"`
	IsStream bool           `json:"stream"`
}
type ollamaGenerateResponse struct {
	Model              string    `json:"model"`
	CreatedAt          time.Time `json:"created_at"`
	Response           string    `json:"response"`
	Done               bool      `json:"done"`
	DoneReason         string    `json:"done_reason"`
	Context            []int     `json:"context"`
	TotalDuration      int64     `json:"total_duration"`
	LoadDuration       int       `json:"load_duration"`
	PromptEvalCount    int       `json:"prompt_eval_count"`
	PromptEvalDuration int       `json:"prompt_eval_duration"`
	EvalCount          int       `json:"eval_count"`
	EvalDuration       int64     `json:"eval_duration"`
}

type ollamaClient struct {
	baseClient
}

func NewOllamaClient()

func (c *ollamaClient) Generate(ctx context.Context, params *GenerateParams) (GenerateResponse, error) {
	url := c.baseUrl + "/api/generate"

	ollamaParams := &ollamaGenerateParams{
		Prompt:   params.Prompt,
		Model:    params.Model,
		Format:   params.Format,
		IsStream: params.IsStreamed,
	}

	jsonData, err := json.Marshal(*ollamaParams)
	if err != nil {
		return GenerateResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return GenerateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GenerateResponse{}, err
	}
	defer res.Body.Close()

	var response GenerateResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return GenerateResponse{}, err
	}

	return response, nil
}
