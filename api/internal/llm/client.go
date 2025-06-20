package llm

import "context"

type GenerateParams struct {
	Model      string
	Prompt     string
	Format     map[string]any
	IsStreamed bool
}

type GenerateResponse struct {
	Content string
}

type LLMClient interface {
	Generate(ctx context.Context, params *GenerateParams) (GenerateResponse, error)
}
