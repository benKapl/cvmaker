package llm

type LLMOffer struct {
	Missions        string
	ExpectedProfile string
}

func ParseOffer(rawOffer string) (LLMOffer, error) {
	return LLMOffer{
		Missions:        "Suck dicks, Make coffee",
		ExpectedProfile: "10 years in backend development",
	}, nil
}

func formatLLMOffer(input string) (output []string, err error) {
	return []string{"one", "two"}, nil
}
