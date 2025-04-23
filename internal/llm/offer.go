package llm

type LLMOffer struct {
	Label                   string
	Organization            string
	OrganizationDescription string
	Missions                string
	Stack                   string
	ExpectedProfile         string
	Miscellaneous           string
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
