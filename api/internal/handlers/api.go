package handlers

import (
	"github.com/benKapl/cvmaker-api/internal/services"
)

type API struct {
	AdminService   *services.AdminService
	AuthService    *services.AuthService
	OfferService   *services.OfferService
	ProfileService *services.ProfileService
	ResumeService  *services.ResumeService
}

func NewAPI(
	adminSrv *services.AdminService,
	authSrv *services.AuthService,
	offerSrv *services.OfferService,
	profileSrv *services.ProfileService,
	resumeSrv *services.ResumeService,
) *API {
	return &API{
		AdminService:   adminSrv,
		AuthService:    authSrv,
		OfferService:   offerSrv,
		ProfileService: profileSrv,
		ResumeService:  resumeSrv,
	}
}
