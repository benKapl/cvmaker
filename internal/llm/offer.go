package llm

import (
	"encoding/json"
	"fmt"
)

type LLMOffer struct {
	Title                   string   `json:"title"`
	Organization            string   `json:"organization"`
	OrganizationDescription string   `json:"organization_description,omitempty"`
	Missions                []string `json:"missions"`
	Stack                   []string `json:"stack,omitempty"`
	ExpectedProfile         []string `json:"expected_profile"`
	Miscellaneous           []string `json:"miscellaneous,omitempty"`
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

// Call the LLM generation endpoint to parse a job offer into a specific format
// Decodes the response into a strongly type LLMOffer and returns it
func (c *Client) ParseOffer(rawOffer string) (LLMOffer, error) {

	prompt := fmt.Sprintf("%s%s%s", offerPromptStart, rawOffer, offerPromptEnd)
	generateResponse, err := c.generate(prompt, offerFormat)
	if err != nil {
		return LLMOffer{}, fmt.Errorf("LLM generation failed: %w", err)
	}

	formattedOffer := generateResponse.Response
	if formattedOffer == "" {
		return LLMOffer{}, fmt.Errorf("LLM Response is empty")
	}

	var offer LLMOffer
	err = json.Unmarshal([]byte(formattedOffer), &offer)
	if err != nil {
		return LLMOffer{}, fmt.Errorf("failed to unmarshal LLM response into LLMOffer: %w. JSON data: %s", err, formattedOffer)
	}

	// Handle missing required values
	if offer.Title == "" {
		offer.Title = "N/A"
	}
	if offer.Organization == "" {
		offer.Organization = "N/A"
	}
	if offer.Missions == nil {
		offer.Missions = []string{}
	}
	if offer.ExpectedProfile == nil {
		offer.ExpectedProfile = []string{}
	}

	return offer, nil
}
