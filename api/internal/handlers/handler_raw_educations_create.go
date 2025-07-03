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

type Education struct {
	Id          uuid.UUID    `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Label       string       `json:"label"`
	School      string       `json:"school"`
	Description string       `json:"description"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     sql.NullTime `json:"end_date"`
	UserID      uuid.UUID    `json:"user_id"`
}

func (a *API) handlerRawEducationsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label       *string    `json:"label"`
		School      *string    `json:"school"`
		Description *string    `json:"description"`
		StartDate   *time.Time `json:"start_date"`
		EndDate     time.Time  `json:"end_date,omitempty"`
	}

	type response struct {
		Success   bool      `json:"success"`
		Education Education `json:"education"`
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
	if params.Label == nil || *params.Label == "" {
		respond.WithError(w, http.StatusBadRequest, "label is a required field", services.ErrMissingField)
		return
	}
	if params.School == nil || *params.School == "" {
		respond.WithError(w, http.StatusBadRequest, "school is a required field", services.ErrMissingField)
		return
	}
	if params.Description == nil || *params.Description == "" {
		respond.WithError(w, http.StatusBadRequest, "description is a required field", services.ErrMissingField)
		return
	}

	rawEducation, err := a.ProfileService.CreateRawEducation(r.Context(), userID, services.CreateRawEducationParams{
		Label:       *params.Label,
		School:      *params.School,
		Description: *params.Description,
		StartDate:   params.StartDate,
		EndDate:     params.EndDate,
	})
	if err != nil {
		if errors.Is(err, services.ErrDuplicateKey) {
			respond.WithError(w, http.StatusBadRequest, "Duplicate key found", err)
			return
		}
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create rawEducation", err)
		return
	}

	respond.WithJSON(w, http.StatusCreated, response{
		Success: true,
		Education: Education{
			Id:          rawEducation.ID,
			CreatedAt:   rawEducation.CreatedAt,
			UpdatedAt:   rawEducation.UpdatedAt,
			Label:       rawEducation.Label,
			School:      rawEducation.School,
			Description: rawEducation.Description,
			StartDate:   rawEducation.StartDate,
			EndDate:     rawEducation.EndDate,
			UserID:      rawEducation.UserID,
		},
	})
}
