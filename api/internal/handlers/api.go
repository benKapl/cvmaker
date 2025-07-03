package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	DB             *database.Queries
	AdminService   *services.AdminService
	AuthService    *services.AuthService
	OfferService   *services.OfferService
	ProfileService *services.ProfileService
}

func NewAPI(db *database.Queries, adminSrv *services.AdminService, authSrv *services.AuthService, offerSrv *services.OfferService, profileSrv *services.ProfileService) *API {
	return &API{
		DB:             db,
		AdminService:   adminSrv,
		AuthService:    authSrv,
		OfferService:   offerSrv,
		ProfileService: profileSrv,
	}
}
