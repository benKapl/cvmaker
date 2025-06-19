package llm

import "context"

type GenerateRequest struct {
	Model      string
	Prompt     string
	Format     map[string]any
	isStreamed bool
}

type GenerateResponse struct {
	Content string
}

type LLMClient interface {
	Generate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error)
}
