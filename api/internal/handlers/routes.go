package handlers

import (
	"log"
	"net/http"
)

func (a *API) RegisterRoutes(mux *http.ServeMux) {

	// Public routes
	mux.HandleFunc("GET /api/healthz", handlerCheckHealth)

	mux.HandleFunc("POST /api/reset", a.handlerReset)
	mux.HandleFunc("POST /api/users", a.handlerUsersCreate)
	mux.HandleFunc("POST /api/login", a.handlerLogin)
	mux.HandleFunc("POST /api/refresh", a.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", a.handlerRevoke)

	// Authenticated routes
	// User history
	mux.Handle("POST /api/raw/hobbies", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerRawHobbiesCreate)))
	mux.Handle("POST /api/raw/stacks", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerRawStacksCreate)))
	mux.Handle("POST /api/raw/educations", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerRawEducationsCreate)))
	mux.Handle("POST /api/raw/experiences", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerRawExperiencesCreate)))
	mux.Handle("POST /api/raw/projects", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerRawProjectsCreate)))
	// Offers management
	mux.Handle("POST /api/offers", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerOffersCreate)))
	// Resume management
	mux.Handle("POST /api/resumes", a.AuthenticateMiddleware(http.HandlerFunc(a.handlerResumesCreate)))

	log.Println("Registered API routes")
}
