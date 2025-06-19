package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/llm"
)

type API struct {
	DB        *database.Queries
	LLMClient llm.LLMClient
	JWTSecret string
	Platform  string
}

func NewAPI(db *database.Queries, llmClient llm.LLMClient, jwtSecret, platform string) *API {
	return &API{
		DB:        db,
		LLMClient: llmClient,
		JWTSecret: jwtSecret,
		Platform:  platform,
	}
}
