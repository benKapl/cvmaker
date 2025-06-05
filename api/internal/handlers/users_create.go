package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/benKapl/cvmaker_api/internal/auth"
	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/respond"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (a *API) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		Email:    strings.ToLower(params.Email),
		Password: hash,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				respond.WithError(w, http.StatusBadRequest, "User already exists", err)
				return
			}
		}
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
