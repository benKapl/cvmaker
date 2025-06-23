package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type JSONSchema struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Schema      map[string]any `json:"schema"`
	Strict      bool           `json:"strict"`
}

type ResponseFormat struct {
	Type       string     `json:"type"`
	JSONSchema JSONSchema `json:"json_schema"`
}
type mistralGenerateParams struct {
	Model          string         `json:"model"`
	Temperature    float64        `json:"temperature,omitempty"`
	TopP           int            `json:"top_p,omitempty"`
	MaxTokens      int            `json:"max_tokens,omitempty"`
	Stream         bool           `json:"stream"`
	Stop           string         `json:"stop,omitempty"`
	RandomSeed     int            `json:"random_seed,omitempty"`
	Messages       []Message      `json:"messages"`
	ResponseFormat ResponseFormat `json:"response_format"`
	Tools          []struct {
		Type     string `json:"type"`
		Function struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Strict      bool   `json:"strict"`
			Parameters  struct {
			} `json:"parameters"`
		} `json:"function"`
	} `json:"tools,omitempty"`
	ToolChoice       string    `json:"tool_choice,omitempty"`
	PresencePenalty  int       `json:"presence_penalty,omitempty"`
	FrequencyPenalty int       `json:"frequency_penalty,omitempty"`
	N                int       `json:"n,omitempty"`
	Prediction       *struct { // Change to pointer type -> zero value will always be ommited
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"prediction,omitempty"`
	ParallelToolCalls bool   `json:"parallel_tool_calls,omitempty"`
	PromptMode        string `json:"prompt_mode,omitempty"`
	SafePrompt        bool   `json:"safe_prompt,omitempty"`
}

type mistralGenerateResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Model  string `json:"model"`
	Usage  struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Created int `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Content   string `json:"content"`
			ToolCalls []struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				Function struct {
					Name      string `json:"name"`
					Arguments struct {
					} `json:"arguments"`
				} `json:"function"`
				Index int `json:"index"`
			} `json:"tool_calls"`
			Prefix bool   `json:"prefix"`
			Role   string `json:"role"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}

type mistralClient struct {
	baseClient
	model string
}

func NewMistralClient(url, apiKey string, timeout time.Duration) *mistralClient {
	baseClient := newBaseClient(url, apiKey, timeout)
	return &mistralClient{
		baseClient: baseClient,
		model:      "mistral-small-latest", // TO BE REFACTOR AS CONFIG PARAMETER !
	}
}

func (c *mistralClient) String() string {
	return fmt.Sprintf("MistralClient (model: %s)", c.model)
}

func (c *mistralClient) Generate(ctx context.Context, params *GenerateParams) (GenerateResponse, error) {
	url := c.baseUrl + "/v1/chat/completions"

	mistralParams := &mistralGenerateParams{
		Model:  c.model,
		Stream: params.Stream,
		Messages: []Message{
			{
				Role:    "user",
				Content: params.Prompt,
			},
		},
		ResponseFormat: ResponseFormat{
			Type: "json_object",
			JSONSchema: JSONSchema{
				Name:   "Formatted JSON Response",
				Schema: params.Format,
				Strict: true,
			},
		},
	}

	jsonData, err := json.Marshal(*mistralParams)
	if err != nil {
		return GenerateResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return GenerateResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return GenerateResponse{}, err
	}
	defer res.Body.Close()

	//	body, err := io.ReadAll(res.Body)
	//	if err != nil {
	//		return GenerateResponse{}, fmt.Errorf("error reading response body: %w", err)
	////////	}
	////////
	//////	log.Println("Response Status:", res.Status)
	//////	log.Println("Response Body:", string(body))
	//////
	var mistralResponse mistralGenerateResponse

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&mistralResponse)
	if err != nil {
		return GenerateResponse{}, err
	}

	response := GenerateResponse{
		Content: mistralResponse.Choices[0].Message.Content,
	}

	return response, nil

}
