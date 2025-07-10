package mocks

import (
	"context"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/stretchr/testify/mock"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	mock.Mock
}

func (m *MockQuerier) CreateOffer(ctx context.Context, arg database.CreateOfferParams) (database.Offer, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.Offer), args.Error(1)
}

func (m *MockQuerier) CreateRawEducation(ctx context.Context, arg database.CreateRawEducationParams) (database.RawEducation, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawEducation), args.Error(1)
}

func (m *MockQuerier) CreateRawExperienceStack(ctx context.Context, arg database.CreateRawExperienceStackParams) (database.CreateRawExperienceStackRow, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.CreateRawExperienceStackRow), args.Error(1)
}

func (m *MockQuerier) CreateRawExperience(ctx context.Context, arg database.CreateRawExperienceParams) (database.RawExperience, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawExperience), args.Error(1)
}

func (m *MockQuerier) CreateRawHobby(ctx context.Context, arg database.CreateRawHobbyParams) (database.RawHobby, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawHobby), args.Error(1)
}

func (m *MockQuerier) CreateRawProjectStack(ctx context.Context, arg database.CreateRawProjectStackParams) (database.CreateRawProjectStackRow, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.CreateRawProjectStackRow), args.Error(1)
}

func (m *MockQuerier) CreateRawProject(ctx context.Context, arg database.CreateRawProjectParams) (database.RawProject, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawProject), args.Error(1)
}

func (m *MockQuerier) CreateRawStack(ctx context.Context, arg database.CreateRawStackParams) (database.RawStack, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawStack), args.Error(1)
}

func (m *MockQuerier) GetRawStackByLabel(ctx context.Context, arg database.GetRawStackByLabelParams) (database.RawStack, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RawStack), args.Error(1)
}

func (m *MockQuerier) CreateRefreshToken(ctx context.Context, arg database.CreateRefreshTokenParams) (database.RefreshToken, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.RefreshToken), args.Error(1)
}

func (m *MockQuerier) GetUserFromRefreshToken(ctx context.Context, token string) (database.User, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockQuerier) RevokeRefreshToken(ctx context.Context, token string) (database.RefreshToken, error) {
	args := m.Called(ctx, token)
	return args.Get(0).(database.RefreshToken), args.Error(1)
}

func (m *MockQuerier) CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockQuerier) DeleteUsers(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockQuerier) GetUser(ctx context.Context, email string) (database.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(database.User), args.Error(1)
}
