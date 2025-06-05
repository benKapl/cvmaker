package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker_api/internal/auth"
	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/respond"
)

func (a *API) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `response:"email"`
		Password string `response:"password"`
	}

	type response struct {
		User
		Token        string
		RefreshToken string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := a.DB.GetUser(r.Context(), params.Email)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Incorrect Email or Password", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.Password)
	if err != nil {
		respond.WithError(w, http.StatusUnauthorized, "Incorrect Email or Password", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, a.JWTSecret, time.Hour)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't make JWT", err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't make refresh token", err)
		return
	}

	_, err = a.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't save refresh token", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}
