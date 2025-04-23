package llm

type LLMConfig struct {
	model      string
	isStreamed bool
}

func NewLLMConfig() LLMConfig {
	return LLMConfig{
		model:      model,
		isStreamed: isStreamed,
	}
}
