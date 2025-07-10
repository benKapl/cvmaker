package prompter_test

//
// type mockLLMClient struct {
// 	response llm.GenerateResponse
// 	err      error
// }
//
// func (m *mockLLMClient) Generate(ctx context.Context, params *llm.GenerateParams) (llm.GenerateResponse, error) {
// 	return m.response, m.err
// }
//
// func (m *mockLLMClient) String() string {
// 	return "mock-llm-client"
// }
//
// func TestParseOffer(t *testing.T) {
// 	validParsedOffer := prompter.ParsedOffer{
// 		Title:           "Software Engineer",
// 		Organization:    "Tech Corp",
// 		Missions:        []string{"Develop applications", "Code review"},
// 		ExpectedProfile: []string{"Go experience", "5+ years"},
// 	}
// 	validJSON, _ := json.Marshal(validParsedOffer)
//
// 	tests := []struct {
// 		name        string
// 		rawOffer    string
// 		mockClient  *mockLLMClient
// 		expected    prompter.ParsedOffer
// 		expectError bool
// 	}{
// 		{
// 			name:     "Valid offer parsing",
// 			rawOffer: "Software Engineer position at Tech Corp. Must have Go experience and 5+ years of development.",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{Content: string(validJSON)},
// 				err:      nil,
// 			},
// 			expected:    validParsedOffer,
// 			expectError: false,
// 		},
// 		{
// 			name:     "LLM generation error",
// 			rawOffer: "Some job offer",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{},
// 				err:      errors.New("LLM service unavailable"),
// 			},
// 			expected:    prompter.ParsedOffer{},
// 			expectError: true,
// 		},
// 		{
// 			name:     "Empty LLM response",
// 			rawOffer: "Some job offer",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{Content: ""},
// 				err:      nil,
// 			},
// 			expected:    prompter.ParsedOffer{},
// 			expectError: true,
// 		},
// 		{
// 			name:     "Invalid JSON response",
// 			rawOffer: "Some job offer",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{Content: "invalid json"},
// 				err:      nil,
// 			},
// 			expected:    prompter.ParsedOffer{},
// 			expectError: true,
// 		},
// 		{
// 			name:     "Partial JSON with missing required fields",
// 			rawOffer: "Some job offer",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{Content: `{"organization": "Company"}`},
// 				err:      nil,
// 			},
// 			expected: prompter.ParsedOffer{
// 				Title:           "N/A",
// 				Organization:    "Company",
// 				Missions:        []string{},
// 				ExpectedProfile: []string{},
// 			},
// 			expectError: false,
// 		},
// 		{
// 			name:     "JSON with null arrays",
// 			rawOffer: "Some job offer",
// 			mockClient: &mockLLMClient{
// 				response: llm.GenerateResponse{Content: `{"title": "Developer", "organization": "Corp", "missions": null, "expected_profile": null}`},
// 				err:      nil,
// 			},
// 			expected: prompter.ParsedOffer{
// 				Title:           "Developer",
// 				Organization:    "Corp",
// 				Missions:        []string{},
// 				ExpectedProfile: []string{},
// 			},
// 			expectError: false,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result, err := prompter.ParseOffer(context.Background(), tt.rawOffer, tt.mockClient)
//
// 			if tt.expectError {
// 				if err == nil {
// 					t.Errorf("ParseOffer() expected error but got none")
// 				}
// 				return
// 			}
//
// 			if err != nil {
// 				t.Errorf("ParseOffer() unexpected error: %v", err)
// 				return
// 			}
//
// 			if result.Title != tt.expected.Title {
// 				t.Errorf("ParseOffer() Title = %v, want %v", result.Title, tt.expected.Title)
// 			}
// 			if result.Organization != tt.expected.Organization {
// 				t.Errorf("ParseOffer() Organization = %v, want %v", result.Organization, tt.expected.Organization)
// 			}
// 			if !stringPtrEqual(result.OrganizationDescription, tt.expected.OrganizationDescription) {
// 				t.Errorf("ParseOffer() OrganizationDescription = %v, want %v", result.OrganizationDescription, tt.expected.OrganizationDescription)
// 			}
// 			if !stringSliceEqual(result.Missions, tt.expected.Missions) {
// 				t.Errorf("ParseOffer() Missions = %v, want %v", result.Missions, tt.expected.Missions)
// 			}
// 			if !stringSliceEqual(result.Stack, tt.expected.Stack) {
// 				t.Errorf("ParseOffer() Stack = %v, want %v", result.Stack, tt.expected.Stack)
// 			}
// 			if !stringSliceEqual(result.ExpectedProfile, tt.expected.ExpectedProfile) {
// 				t.Errorf("ParseOffer() ExpectedProfile = %v, want %v", result.ExpectedProfile, tt.expected.ExpectedProfile)
// 			}
// 			if !stringSliceEqual(result.Miscellaneous, tt.expected.Miscellaneous) {
// 				t.Errorf("ParseOffer() Miscellaneous = %v, want %v", result.Miscellaneous, tt.expected.Miscellaneous)
// 			}
// 		})
// 	}
// }
//
// // Helper functions
// func stringPtrEqual(a, b *string) bool {
// 	if a == nil && b == nil {
// 		return true
// 	}
// 	if a == nil || b == nil {
// 		return false
// 	}
// 	return *a == *b
// }
//
// func stringSliceEqual(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := range a {
// 		if a[i] != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
