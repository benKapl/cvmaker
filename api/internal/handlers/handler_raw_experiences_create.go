package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker-api/internal/respond"
	"github.com/benKapl/cvmaker-api/internal/services"
	"github.com/google/uuid"
)

type Experience struct {
	Id           uuid.UUID    `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Title        string       `json:"title"`
	Organization string       `json:"organization"`
	Description  string       `json:"description"`
	Stacks       []Stack      `json:"stacks,omitempty"`
	StartDate    time.Time    `json:"start_date"`
	EndDate      sql.NullTime `json:"end_date"`
	UserID       uuid.UUID    `json:"user_id"`
}

func (a *API) handlerRawExperiencesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title        *string    `json:"title"`
		Organization *string    `json:"organization"`
		Description  *string    `json:"description"`
		Stacks       []string   `json:"stacks,omitempty"`
		StartDate    *time.Time `json:"start_date"`
		EndDate      time.Time  `json:"end_date,omitempty"`
	}

	type response struct {
		Success    bool       `json:"success"`
		Experience Experience `json:"experience"`
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
	if params.StartDate == nil {
		respond.WithError(w, http.StatusBadRequest, "start_date is a required field", services.ErrMissingField)
		return
	}
	if params.Title == nil || *params.Title == "" {
		respond.WithError(w, http.StatusBadRequest, "title is a required field", services.ErrMissingField)
		return
	}
	if params.Organization == nil || *params.Organization == "" {
		respond.WithError(w, http.StatusBadRequest, "organization is a required field", services.ErrMissingField)
		return
	}
	if params.Description == nil || *params.Description == "" {
		respond.WithError(w, http.StatusBadRequest, "description is a required field", services.ErrMissingField)
		return
	}

	rawExperience, rawStacks, err := a.ProfileService.CreateRawExperience(r.Context(), userID, services.CreateRawExperienceParams{
		Title:        *params.Title,
		Organization: *params.Organization,
		Description:  *params.Description,
		Stacks:       params.Stacks,
		StartDate:    params.StartDate,
		EndDate:      params.EndDate,
	})
	if err != nil {
		if errors.Is(err, services.ErrDuplicateKey) {
			respond.WithError(w, http.StatusConflict, "Duplicate key found", err)
			return
		}
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create rawExperience", err)
		return
	}

	var stacks []Stack
	for _, rawStack := range rawStacks {
		stacks = append(stacks, Stack{
			Id:        rawStack.ID,
			CreatedAt: rawStack.CreatedAt,
			UpdatedAt: rawStack.UpdatedAt,
			Label:     rawStack.Label,
			UserID:    rawStack.UserID,
		})
	}

	respond.WithJSON(w, http.StatusCreated, response{
		Success: true,
		Experience: Experience{
			Id:           rawExperience.ID,
			CreatedAt:    rawExperience.CreatedAt,
			UpdatedAt:    rawExperience.UpdatedAt,
			Title:        rawExperience.Title,
			Organization: rawExperience.Organization,
			Description:  rawExperience.Description,
			Stacks:       stacks,
			StartDate:    rawExperience.StartDate,
			EndDate:      rawExperience.EndDate,
			UserID:       rawExperience.UserID,
		},
	})
}
