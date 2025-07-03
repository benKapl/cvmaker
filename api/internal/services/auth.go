package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/benKapl/cvmaker-api/internal/auth"
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type AuthService struct {
	DB        *database.Queries
	JWTSecret string
}

func NewAuthService(db *database.Queries, jwtSecret string) *AuthService {
	return &AuthService{
		DB:        db,
		JWTSecret: jwtSecret,
	}
}

type LoginResponse struct {
	User         database.User
	AccessToken  string
	RefreshToken string
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateKey       = errors.New("element already exists")
)

// Create user in database
func (s *AuthService) CreateUser(ctx context.Context, email, password string) (database.User, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return database.User{}, fmt.Errorf("couldn't hash password: %w", err)
	}

	user, err := s.DB.CreateUser(ctx, database.CreateUserParams{
		Email:    strings.ToLower(email),
		Password: hash,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.User{}, ErrDuplicateKey
			}
		}
		return database.User{}, fmt.Errorf("couldn't create user: %w", err)
	}

	return user, nil
}

// Authenticate user
func (s *AuthService) Login(ctx context.Context, email, password string) (LoginResponse, error) {

	user, err := s.DB.GetUser(ctx, email)
	if err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	err = auth.CheckPasswordHash(password, user.Password)
	if err != nil {
		return LoginResponse{}, ErrInvalidCredentials
	}

	accessToken, err := auth.MakeJWT(user.ID, s.JWTSecret, time.Hour)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("couldn't make JWT: %w", err)
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("couldn't make refresh token: %w", err)
	}

	_, err = s.DB.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		return LoginResponse{}, fmt.Errorf("couldn't save refresh token: %w", err)
	}

	return LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) ValidateJWT(tokenString string) (uuid.UUID, error) {
	return auth.ValidateJWT(tokenString, s.JWTSecret)
}

func (s *AuthService) RefreshJWT(ctx context.Context, refreshToken string, expiresIn time.Duration) (string, error) {
	user, err := s.DB.GetUserFromRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("couldn't get user for refresh token: %w", err)
	}

	accessToken, err := auth.MakeJWT(
		user.ID,
		s.JWTSecret,
		time.Hour,
	)
	if err != nil {
		return "", fmt.Errorf("couldn't validate JWT: %w", err)
	}

	return accessToken, nil

}
