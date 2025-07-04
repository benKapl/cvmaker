package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker-api/internal/respond"
	"github.com/benKapl/cvmaker-api/internal/services"
	"github.com/google/uuid"
)

type Hobby struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
	UserID    uuid.UUID `json:"user_id"`
}

func (a *API) handlerRawHobbiesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label *string `json:"label"`
	}

	type response struct {
		Success bool  `json:"success"`
		Hobby   Hobby `json:"hobby"`
	}

	userID, ok := r.Context().Value(userIDKey).(uuid.UUID)
	if !ok {
		respond.WithError(w, http.StatusInternalServerError, "Could not get userID, from Context", nil)
		return
	}

	var params parameters

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&params)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	if params.Label == nil || *params.Label == "" {
		respond.WithError(w, http.StatusBadRequest, "label is a required field", services.ErrMissingField)
		return
	}

	rawHobby, err := a.ProfileService.CreateRawHobby(r.Context(), userID, services.CreateRawHobbyParams{
		Label: *params.Label,
	})
	if err != nil {
		if errors.Is(err, services.ErrDuplicateKey) {
			respond.WithError(w, http.StatusBadRequest, "Duplicate key found", err)
			return
		}
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create rawHobby", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		Success: true,
		Hobby: Hobby{
			Id:        rawHobby.ID,
			CreatedAt: rawHobby.CreatedAt,
			UpdatedAt: rawHobby.UpdatedAt,
			Label:     rawHobby.Label,
			UserID:    rawHobby.UserID,
		},
	})
}
