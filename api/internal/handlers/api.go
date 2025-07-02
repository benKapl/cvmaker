package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/llm"
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	DB           *database.Queries
	LLMClient    llm.LLMClient
	JWTSecret    string
	Platform     string
	OfferService *services.OfferService
}

func NewAPI(db *database.Queries, llmClient llm.LLMClient, jwtSecret, platform string, offerSrv *services.OfferService) *API {
	return &API{
		DB:           db,
		LLMClient:    llmClient,
		JWTSecret:    jwtSecret,
		Platform:     platform,
		OfferService: offerSrv,
	}
}
