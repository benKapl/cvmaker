package handlers

import (
	"net/http"
	"time"

	"github.com/benKapl/cvmaker-api/internal/auth"
	"github.com/benKapl/cvmaker-api/internal/respond"
)

func (a *API) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't find token", err)
		return
	}

	user, err := a.DB.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't get user for refresh token", err)
		return
	}

	accessToken, err := auth.MakeJWT(
		user.ID,
		a.JWTSecret,
		time.Hour,
	)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	respond.WithJSON(w, http.StatusOK, response{
		Token: accessToken,
	})
}

func (a *API) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't find token", err)
		return
	}

	_, err = a.DB.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't revoke session", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
