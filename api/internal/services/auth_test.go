
package services

import (
	"context"
	"testing"
	"time"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/mocks"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)



func TestAuthService_CreateUser(t *testing.T) {
	ctx := context.Background()
	mockDb := new(mocks.MockQuerier)
	authService := NewAuthService(mockDb, "test_secret")

	email := "test@example.com"
	password := "password"

	t.Run("successful user creation", func(t *testing.T) {
		expectedUser := database.User{
			ID:        uuid.New(),
			Email:     email,
			Password:  "hashed_password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockDb.On("CreateUser", ctx, mock.AnythingOfType("database.CreateUserParams")).Return(expectedUser, nil).Once()

		user, err := authService.CreateUser(ctx, email, password)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockDb.AssertExpectations(t)
	})

	t.Run("duplicate user", func(t *testing.T) {
		pqErr := &pq.Error{Code: "23505"}
		mockDb.On("CreateUser", ctx, mock.AnythingOfType("database.CreateUserParams")).Return(database.User{}, pqErr).Once()

		_, err := authService.CreateUser(ctx, email, password)

		assert.ErrorIs(t, err, ErrDuplicateKey)
		mockDb.AssertExpectations(t)
	})
}
