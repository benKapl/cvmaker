package prompter

import (
	"context"

	"github.com/benKapl/cvmaker-api/internal/llm"
)

type Prompter interface {
	ParseOffer(ctx context.Context, rawOffer string) (ParsedOffer, error)
}

type DefaultPrompter struct {
	LLMClient llm.LLMClient
}

func NewDefaultPrompter(llmClient llm.LLMClient) *DefaultPrompter {
	return &DefaultPrompter{
		LLMClient: llmClient,
	}
}

var _ Prompter = (*DefaultPrompter)(nil)
