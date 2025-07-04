package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	AdminService   *services.AdminService
	AuthService    *services.AuthService
	OfferService   *services.OfferService
	ProfileService *services.ProfileService
}

func NewAPI(adminSrv *services.AdminService, authSrv *services.AuthService, offerSrv *services.OfferService, profileSrv *services.ProfileService) *API {
	return &API{
		AdminService:   adminSrv,
		AuthService:    authSrv,
		OfferService:   offerSrv,
		ProfileService: profileSrv,
	}
}
