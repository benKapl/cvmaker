package mocks

import (
	"context"

	"github.com/benKapl/cvmaker-api/internal/prompter"
	"github.com/stretchr/testify/mock"
)

// MockPrompter is a mock of the Prompter interface
type MockPrompter struct {
	mock.Mock
}

// ParseOffer implements prompter.Prompter
func (m *MockPrompter) ParseOffer(ctx context.Context, rawOffer string) (prompter.ParsedOffer, error) {
	args := m.Called(ctx, rawOffer)
	return args.Get(0).(prompter.ParsedOffer), args.Error(1)
}
