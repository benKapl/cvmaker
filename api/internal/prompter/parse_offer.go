package prompter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/benKapl/cvmaker-api/internal/llm"
)

type ParsedOffer struct {
	Title                   string   `json:"title"`
	Organization            string   `json:"organization"`
	OrganizationDescription *string  `json:"organization_description,omitempty"`
	Missions                []string `json:"missions"`
	Stack                   []string `json:"stack,omitempty"`
	ExpectedProfile         []string `json:"expected_profile"`
	Miscellaneous           []string `json:"miscellaneous,omitempty"`
}

var (
	offerFormat = map[string]any{
		"type": "object",
		"properties": map[string]any{
			"title":                    map[string]string{"type": "string"},
			"organization":             map[string]string{"type": "string"},
			"organization_description": map[string]string{"type": "string"},
			"missions": map[string]any{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"stack": map[string]any{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"expected_profile": map[string]any{
				"type":  "array",
				"items": map[string]string{"type": "string"},
			},
			"miscellaneous": map[string]any{
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
	offerPromptStart = "Extract the following information from the job offer provided below and return it as a JSON object:\n\n- title (required)\n- organization (required) - the name of the postee\n- organization_description\n- missions (required) - a list of tasks and responsibilities.\n- stack - a list of technologies.\n- expected_profile (required) - a list of required skills, experience, and qualifications. Break down sentences into separate points in the list.\n- miscellaneous - a list of all other information. Each piece of information should be a separate string in the list. For example: [\"Location: Paris, France\", \"Salary: 45000 euros per year\", \"Remote work: Accepted\"]\n\nJob Offer:\n\"\"\"\n"
	offerPromptEnd   = "\n\"\"\"\n\nJSON Output:"
)

func (p *DefaultPrompter) ParseOffer(ctx context.Context, rawOffer string) (ParsedOffer, error) {
	prompt := fmt.Sprintf("%s%s%s", offerPromptStart, rawOffer, offerPromptEnd)
	params := &llm.GenerateParams{
		Prompt: prompt,
		Format: offerFormat,
		Stream: false,
	}

	response, err := p.LLMClient.Generate(ctx, params)
	if err != nil {
		return ParsedOffer{}, fmt.Errorf("LLM generation failed: %w", err)
	}

	parsedOffer := response.Content
	if parsedOffer == "" {
		return ParsedOffer{}, fmt.Errorf("LLM response is empty: %w", err)
	}

	var offer ParsedOffer
	err = json.Unmarshal([]byte(parsedOffer), &offer)
	if err != nil {
		return ParsedOffer{}, fmt.Errorf("failed to unmarshal LLM response into ParsedOffer: %w. JSON data: %s", err, parsedOffer)
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
