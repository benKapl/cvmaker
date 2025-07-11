package services

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/benKapl/cvmaker-api/internal/auth"
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
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")
		mockDb.On("CreateUser", ctx, mock.AnythingOfType("database.CreateUserParams")).Return(database.User{}, pqErr).Once()

		_, err := authService.CreateUser(ctx, email, password)

		assert.ErrorIs(t, err, ErrDuplicateKey)
		mockDb.AssertExpectations(t)
	})
}

func TestAuthService_Login(t *testing.T) {
	ctx := context.Background()
	email := "test@example.com"
	password := "password"
	hashedPassword, _ := auth.HashPassword(password)
	userID := uuid.New()

	t.Run("successful login", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")

		user := database.User{ID: userID, Email: email, Password: hashedPassword}
		mockDb.On("GetUser", ctx, email).Return(user, nil).Once()
		mockDb.On("CreateRefreshToken", ctx, mock.AnythingOfType("database.CreateRefreshTokenParams")).Return(database.RefreshToken{}, nil).Once()

		loginResp, err := authService.Login(ctx, email, password)

		assert.NoError(t, err)
		assert.Equal(t, user, loginResp.User)
		assert.NotEmpty(t, loginResp.AccessToken)
		assert.NotEmpty(t, loginResp.RefreshToken)
		mockDb.AssertExpectations(t)
	})

	t.Run("invalid credentials - user not found", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")

		mockDb.On("GetUser", ctx, email).Return(database.User{}, errors.New("not found")).Once()

		_, err := authService.Login(ctx, email, password)

		assert.ErrorIs(t, err, ErrInvalidCredentials)
		mockDb.AssertExpectations(t)
	})

	t.Run("invalid credentials - wrong password", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")

		user := database.User{ID: userID, Email: email, Password: "wrong_hashed_password"}
		mockDb.On("GetUser", ctx, email).Return(user, nil).Once()

		_, err := authService.Login(ctx, email, "wrong_password")

		assert.ErrorIs(t, err, ErrInvalidCredentials)
		mockDb.AssertExpectations(t)
	})

	t.Run("error creating refresh token", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")

		user := database.User{ID: userID, Email: email, Password: hashedPassword}
		mockDb.On("GetUser", ctx, email).Return(user, nil).Once()
		mockDb.On("CreateRefreshToken", ctx, mock.AnythingOfType("database.CreateRefreshTokenParams")).Return(database.RefreshToken{}, errors.New("db error")).Once()

		_, err := authService.Login(ctx, email, password)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "couldn't save refresh token")
		mockDb.AssertExpectations(t)
	})
}

func TestAuthService_RefreshJWT(t *testing.T) {
	ctx := context.Background()
	mockDb := new(mocks.MockQuerier)
	authService := NewAuthService(mockDb, "test_secret")
	refreshToken := "valid-refresh-token"
	userID := uuid.New()

	t.Run("successful refresh", func(t *testing.T) {
		user := database.User{ID: userID, Email: "test@example.com"}
		mockDb.On("GetUserFromRefreshToken", ctx, refreshToken).Return(user, nil).Once()

		accessToken, err := authService.RefreshJWT(ctx, refreshToken, time.Hour)

		assert.NoError(t, err)
		assert.NotEmpty(t, accessToken)

		// Validate the new token to ensure it's a valid JWT for the correct user
		validatedUserID, err := auth.ValidateJWT(accessToken, "test_secret")
		assert.NoError(t, err)
		assert.Equal(t, userID, validatedUserID)

		mockDb.AssertExpectations(t)
	})

	t.Run("error getting user from refresh token", func(t *testing.T) {
		// Reset the mock for this specific sub-test
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")
		dbErr := errors.New("invalid refresh token")
		mockDb.On("GetUserFromRefreshToken", ctx, refreshToken).Return(database.User{}, dbErr).Once()

		accessToken, err := authService.RefreshJWT(ctx, refreshToken, time.Hour)

		assert.Error(t, err)
		assert.Empty(t, accessToken)
		assert.Contains(t, err.Error(), "couldn't get user for refresh token")
		mockDb.AssertExpectations(t)
	})
}

func TestAuthService_RevokeRefreshToken(t *testing.T) {
	ctx := context.Background()
	mockDb := new(mocks.MockQuerier)
	authService := NewAuthService(mockDb, "test_secret")
	token := "some-refresh-token"

	t.Run("successful revocation", func(t *testing.T) {
		expectedToken := database.RefreshToken{
			Token:     token,
			UserID:    uuid.New(),
			ExpiresAt: time.Now().Add(time.Hour),
			RevokedAt: sql.NullTime{Time: time.Now(), Valid: true},
		}
		mockDb.On("RevokeRefreshToken", ctx, token).Return(expectedToken, nil).Once()

		result, err := authService.RevokeRefreshToken(ctx, token)

		assert.NoError(t, err)
		assert.Equal(t, expectedToken, result)
		mockDb.AssertExpectations(t)
	})

	t.Run("db error on revocation", func(t *testing.T) {
		// Reset the mock for this specific sub-test
		mockDb := new(mocks.MockQuerier)
		authService := NewAuthService(mockDb, "test_secret")
		dbErr := errors.New("database error")
		mockDb.On("RevokeRefreshToken", ctx, token).Return(database.RefreshToken{}, dbErr).Once()

		_, err := authService.RevokeRefreshToken(ctx, token)

		assert.ErrorIs(t, err, dbErr)
		mockDb.AssertExpectations(t)
	})
}
