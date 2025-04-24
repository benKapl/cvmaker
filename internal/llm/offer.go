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

var (
	offerFormat = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"title":                    map[string]string{"type": "string"},
			"organization":             map[string]string{"type": "string"},
			"organization_description": map[string]string{"type": "string"},
			"missions": map[string]interface{}{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"stack": map[string]interface{}{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"expected_profile": map[string]interface{}{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"miscellaneous": map[string]interface{}{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
		},
		"required": []string{
			"title",
			"organization",
			"missions",
			"expected_profile",
		},
	}
	offerPromptStart = "Extract the following information from the job offer provided below and return it as a JSON object:\n\n- title (required)\n- organization (required) - the name of the postee\n- organization_description\n- missions (required)\n- stack\n- expected_profile (required)\n- miscellaneous (all other relevant info)\n\nEnsure the output is a valid JSON object matching the specified structure. If a field is not found and is not required, omit it. Place any extra information not fitting the specific fields into the miscellaneous field.\n\nJob Offer:\n\"\"\"\n"
	offerPromptEnd   = "\n\"\"\"\n\nJSON Output:"
)

func ParseOffer(rawOffer string) (LLMOffer, error) {
	// REVIEW STRUCT MODEL !!!

	// prompt := fmt.Sprintf("%s%s%s", offerPromptStart, rawOffer, offerPromptEnd)
	return LLMOffer{
		Missions:        "Suck dicks, Make coffee",
		ExpectedProfile: "10 years in backend development",
	}, nil
}

func formatLLMOffer(input string) (output []string, err error) {
	return []string{"one", "two"}, nil
}
