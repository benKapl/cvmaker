package llm

import (
	"context"
	"fmt"
	"time"
)

type mistralGenerateParams struct {
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature"`
	TopP        int     `json:"top_p"`
	MaxTokens   int     `json:"max_tokens"`
	Stream      bool    `json:"stream"`
	Stop        string  `json:"stop"`
	RandomSeed  int     `json:"random_seed"`
	Messages    []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	ResponseFormat struct {
		Type       string `json:"type"`
		JSONSchema struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Schema      map[string]any `json:"schema"`
			Strict      bool           `json:"strict"`
		} `json:"json_schema"`
	} `json:"response_format"`
	Tools []struct {
		Type     string `json:"type"`
		Function struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Strict      bool   `json:"strict"`
			Parameters  struct {
			} `json:"parameters"`
		} `json:"function"`
	} `json:"tools"`
	ToolChoice       string `json:"tool_choice"`
	PresencePenalty  int    `json:"presence_penalty"`
	FrequencyPenalty int    `json:"frequency_penalty"`
	N                int    `json:"n"`
	Prediction       struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"prediction"`
	ParallelToolCalls bool   `json:"parallel_tool_calls"`
	PromptMode        string `json:"prompt_mode"`
	SafePrompt        bool   `json:"safe_prompt"`
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

func (c *mistralClient) Generate(ctx context.Context, params *GenerateParams) (GenerateResponse, error)
