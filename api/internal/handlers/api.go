package handlers

import (
	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/llm"
)

type API struct {
	DB        *database.Queries
	LLMClient llm.Client
	JWTSecret string
	Platform  string
}

func NewAPI(db *database.Queries, llmClient llm.Client, jwtSecret, platform string) *API {
	return &API{
		DB:        db,
		LLMClient: llmClient,
		JWTSecret: jwtSecret,
		Platform:  platform,
	}
}
