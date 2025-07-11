package mocks

import (
	"context"

	"github.com/benKapl/cvmaker-api/internal/llm"
	"github.com/stretchr/testify/mock"
)

// MockLLMClient is a mock of the LLMClient interface
type MockLLMClient struct {
	mock.Mock
}

// Generate implements llm.LLMClient
func (m *MockLLMClient) Generate(ctx context.Context, params *llm.GenerateParams) (llm.GenerateResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(llm.GenerateResponse), args.Error(1)
}

// String implements llm.LLMClient
func (m *MockLLMClient) String() string {
	args := m.Called()
	return args.String(0)
}
