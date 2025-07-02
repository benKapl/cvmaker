package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/benKapl/cvmaker-api/internal/auth"
	"github.com/benKapl/cvmaker-api/internal/database"
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

var ErrInvalidCredentials = errors.New("invalid credentials")

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
