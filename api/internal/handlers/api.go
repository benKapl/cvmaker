package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	DB           *database.Queries
	Platform     string
	AuthService  *services.AuthService
	OfferService *services.OfferService
}

func NewAPI(db *database.Queries, platform string, authSrv *services.AuthService, offerSrv *services.OfferService) *API {
	return &API{
		DB:           db,
		Platform:     platform,
		AuthService:  authSrv,
		OfferService: offerSrv,
	}
}
