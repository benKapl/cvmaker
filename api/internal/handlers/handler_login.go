package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/benKapl/cvmaker-api/internal/respond"
	"github.com/benKapl/cvmaker-api/internal/services"
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

	loginResponse, err := a.AuthService.Login(r.Context(), params.Email, params.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			respond.WithError(w, http.StatusUnauthorized, "Incorrect Email or Password", err)
			return
		}
		respond.WithError(w, http.StatusInternalServerError, "Failed to log in", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        loginResponse.User.ID,
			CreatedAt: loginResponse.User.CreatedAt,
			UpdatedAt: loginResponse.User.UpdatedAt,
			Email:     loginResponse.User.Email,
		},
		Token:        loginResponse.AccessToken,
		RefreshToken: loginResponse.RefreshToken,
	})
}
