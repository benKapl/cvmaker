package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ollamaGenerateParams struct {
	Model  string         `json:"model"`
	Prompt string         `json:"prompt"`
	Format map[string]any `json:"format,omitempty"`
	Stream bool           `json:"stream"`
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

type OllamaClient struct {
	baseClient
	model string
}

func NewOllamaClient(url string, timeout time.Duration) *OllamaClient {
	baseClient := newBaseClient(url, "", timeout)
	return &OllamaClient{
		baseClient: baseClient,
		model:      "mistral", // TO BE REFACTOR AS CONFIG PARAMETER !
	}
}
func (c *OllamaClient) String() string {
	return fmt.Sprintf("OllamaClient (model: %s)", c.model)
}

func (c *OllamaClient) Generate(ctx context.Context, params *GenerateParams) (GenerateResponse, error) {
	url := c.baseUrl + "/api/generate"

	ollamaParams := &ollamaGenerateParams{
		Model:  c.model,
		Prompt: params.Prompt,
		Format: params.Format,
		Stream: params.Stream,
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

	var ollamaResponse ollamaGenerateResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&ollamaResponse)
	if err != nil {
		return GenerateResponse{}, err
	}

	response := GenerateResponse{
		Content: ollamaResponse.Response,
	}

	return response, nil
}
