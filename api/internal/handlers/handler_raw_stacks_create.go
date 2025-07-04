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

type Stack struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
	UserID    uuid.UUID `json:"user_id"`
}

func (a *API) handlerRawStacksCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label *string `json:"label"`
	}

	type response struct {
		Success bool  `json:"success"`
		Stack   Stack `json:"stack"`
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

	rawStack, err := a.ProfileService.CreateRawStack(r.Context(), userID, services.CreateRawStackParams{
		Label: *params.Label,
	})
	if err != nil {
		if errors.Is(err, services.ErrDuplicateKey) {
			respond.WithError(w, http.StatusBadRequest, "Duplicate key found", err)
			return
		}
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create rawStack", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		Success: true,
		Stack: Stack{
			Id:        rawStack.ID,
			CreatedAt: rawStack.CreatedAt,
			UpdatedAt: rawStack.UpdatedAt,
			Label:     rawStack.Label,
			UserID:    rawStack.UserID,
		},
	})
}
