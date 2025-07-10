package prompter

import "github.com/benKapl/cvmaker-api/internal/llm"

type Prompter struct {
	LLMClient llm.LLMClient
}

func New(llmClient llm.LLMClient) *Prompter {
	return &Prompter{
		LLMClient: llmClient,
	}
}
