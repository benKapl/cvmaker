package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type GenerateParams struct {
	Model  string         `json:"model"`
	Prompt string         `json:"prompt"`
	Format map[string]any `json:"format,omitempty"`
	Stream bool           `json:"stream"`
}
type GenerateResponse struct {
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

func (lc *LLMClient) generate(prompt string, format map[string]any) (GenerateResponse, error) {
	url := lc.baseUrl + "/api/generate"

	params := GenerateParams{
		Prompt: prompt,
		Model:  lc.llmConfig.model,
		Format: format,
		Stream: lc.llmConfig.isStreamed,
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return GenerateResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return GenerateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := lc.httpClient.Do(req)
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
