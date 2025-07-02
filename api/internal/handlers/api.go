package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	DB           *database.Queries
	JWTSecret    string
	Platform     string
	AuthService  *services.AuthService
	OfferService *services.OfferService
}

func NewAPI(db *database.Queries, jwtSecret, platform string, authSrv *services.AuthService, offerSrv *services.OfferService) *API {
	return &API{
		DB:           db,
		JWTSecret:    jwtSecret,
		Platform:     platform,
		AuthService:  authSrv,
		OfferService: offerSrv,
	}
}
