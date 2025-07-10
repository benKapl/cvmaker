package prompter_test

import (
	"context"
	"errors"
	"testing"

	"github.com/benKapl/cvmaker-api/internal/llm"
	"github.com/benKapl/cvmaker-api/internal/mocks"
	"github.com/benKapl/cvmaker-api/internal/prompter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParseOffer_Success(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Software Engineer at TechCorp\nDevelop web applications using React and Node.js\nRequires 3+ years experience"

	expectedResponse := llm.GenerateResponse{
		Content: `{
			"title": "Software Engineer",
			"organization": "TechCorp",
			"organization_description": "Leading tech company",
			"missions": ["Develop web applications", "Maintain existing systems"],
			"stack": ["React", "Node.js"],
			"expected_profile": ["3+ years experience", "JavaScript expertise"],
			"miscellaneous": ["Remote work available"]
		}`,
	}

	mockClient.On("Generate", mock.Anything, mock.MatchedBy(func(params *llm.GenerateParams) bool {
		return params.Prompt != "" && params.Format != nil && !params.Stream
	})).Return(expectedResponse, nil)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.NoError(t, err)
	assert.Equal(t, "Software Engineer", result.Title)
	assert.Equal(t, "TechCorp", result.Organization)
	assert.Equal(t, "Leading tech company", *result.OrganizationDescription)
	assert.Equal(t, []string{"Develop web applications", "Maintain existing systems"}, result.Missions)
	assert.Equal(t, []string{"React", "Node.js"}, result.Stack)
	assert.Equal(t, []string{"3+ years experience", "JavaScript expertise"}, result.ExpectedProfile)
	assert.Equal(t, []string{"Remote work available"}, result.Miscellaneous)

	mockClient.AssertExpectations(t)
}

func TestParseOffer_LLMError(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Test job offer"
	expectedError := errors.New("LLM service unavailable")

	mockClient.On("Generate", mock.Anything, mock.AnythingOfType("*llm.GenerateParams")).Return(llm.GenerateResponse{}, expectedError)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "LLM generation failed")
	assert.Equal(t, prompter.ParsedOffer{}, result)

	mockClient.AssertExpectations(t)
}

func TestParseOffer_EmptyResponse(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Test job offer"
	emptyResponse := llm.GenerateResponse{Content: ""}

	mockClient.On("Generate", mock.Anything, mock.AnythingOfType("*llm.GenerateParams")).Return(emptyResponse, nil)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "LLM response is empty")
	assert.Equal(t, prompter.ParsedOffer{}, result)

	mockClient.AssertExpectations(t)
}

func TestParseOffer_InvalidJSON(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Test job offer"
	invalidJSONResponse := llm.GenerateResponse{Content: "invalid json response"}

	mockClient.On("Generate", mock.Anything, mock.AnythingOfType("*llm.GenerateParams")).Return(invalidJSONResponse, nil)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal LLM response into ParsedOffer")
	assert.Contains(t, err.Error(), "invalid json response")
	assert.Equal(t, prompter.ParsedOffer{}, result)

	mockClient.AssertExpectations(t)
}

func TestParseOffer_MissingRequiredFields(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Test job offer"

	partialResponse := llm.GenerateResponse{
		Content: `{
			"organization_description": "Some company",
			"stack": ["Go", "PostgreSQL"]
		}`,
	}

	mockClient.On("Generate", mock.Anything, mock.AnythingOfType("*llm.GenerateParams")).Return(partialResponse, nil)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.NoError(t, err)
	assert.Equal(t, "N/A", result.Title)
	assert.Equal(t, "N/A", result.Organization)
	assert.Equal(t, "Some company", *result.OrganizationDescription)
	assert.Equal(t, []string{"Go", "PostgreSQL"}, result.Stack)
	assert.Equal(t, []string{}, result.Missions)
	assert.Equal(t, []string{}, result.ExpectedProfile)
	assert.Nil(t, result.Miscellaneous)

	mockClient.AssertExpectations(t)
}

func TestParseOffer_ValidMinimalResponse(t *testing.T) {
	mockClient := &mocks.MockLLMClient{}
	p := prompter.NewDefaultPrompter(mockClient)

	rawOffer := "Test job offer"

	minimalResponse := llm.GenerateResponse{
		Content: `{
			"title": "Developer",
			"organization": "TestCorp",
			"missions": ["Code", "Test"],
			"expected_profile": ["Experience required"]
		}`,
	}

	mockClient.On("Generate", mock.Anything, mock.AnythingOfType("*llm.GenerateParams")).Return(minimalResponse, nil)

	result, err := p.ParseOffer(context.Background(), rawOffer)

	assert.NoError(t, err)
	assert.Equal(t, "Developer", result.Title)
	assert.Equal(t, "TestCorp", result.Organization)
	assert.Nil(t, result.OrganizationDescription)
	assert.Equal(t, []string{"Code", "Test"}, result.Missions)
	assert.Nil(t, result.Stack)
	assert.Equal(t, []string{"Experience required"}, result.ExpectedProfile)
	assert.Nil(t, result.Miscellaneous)

	mockClient.AssertExpectations(t)
}

// func TestParseOffer_PromptConstruction(t *testing.T) {
// 	mockClient := &mocks.MockLLMClient{}
// 	prompter := prompter.NewDefaultPrompter(mockClient)
//
// 	rawOffer := "Test job offer content"
//
// 	response := llm.GenerateResponse{
// 		Content: `{
// 			"title": "Test Title",
// 			"organization": "Test Org",
// 			"missions": ["Test Mission"],
// 			"expected_profile": ["Test Profile"]
// 		}`,
// 	}
//
// 	mockClient.On("Generate", mock.Anything, mock.MatchedBy(func(params *llm.GenerateParams) bool {
// 		expectedPrompt := offerPromptStart + rawOffer + offerPromptEnd
// 		return params.Prompt == expectedPrompt &&
// 			params.Format == offerFormat &&
// 			params.Stream == false
// 	})).Return(response, nil)
//
// 	_, err := prompter.ParseOffer(context.Background(), rawOffer)
//
// 	assert.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }
