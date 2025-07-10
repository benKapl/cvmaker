package database

import (
	"context"
)

type Querier interface {
	CreateOffer(ctx context.Context, arg CreateOfferParams) (Offer, error)
	CreateRawEducation(ctx context.Context, arg CreateRawEducationParams) (RawEducation, error)
	CreateRawExperienceStack(ctx context.Context, arg CreateRawExperienceStackParams) (CreateRawExperienceStackRow, error)
	CreateRawExperience(ctx context.Context, arg CreateRawExperienceParams) (RawExperience, error)
	CreateRawHobby(ctx context.Context, arg CreateRawHobbyParams) (RawHobby, error)
	CreateRawProjectStack(ctx context.Context, arg CreateRawProjectStackParams) (CreateRawProjectStackRow, error)
	CreateRawProject(ctx context.Context, arg CreateRawProjectParams) (RawProject, error)
	CreateRawStack(ctx context.Context, arg CreateRawStackParams) (RawStack, error)
	GetRawStackByLabel(ctx context.Context, arg GetRawStackByLabelParams) (RawStack, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	GetUserFromRefreshToken(ctx context.Context, token string) (User, error)
	RevokeRefreshToken(ctx context.Context, token string) (RefreshToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUsers(ctx context.Context) error
	GetUser(ctx context.Context, email string) (User, error)
}

// Check at compile time if *Queries implement the Querier interface
var _ Querier = (*Queries)(nil)
