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
		Label       string    `json:"label"`
		School      string    `json:"school"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
	}

	type response struct {
		Success   bool `json:"success"`
		Education Education
	}

	userID, ok := r.Context().Value(userIDKey).(uuid.UUID)
	if !ok {
		respond.WithError(w, http.StatusInternalServerError, "Could not get userID, from Context", nil)
		return
	}

	var params parameters

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	rawEducation, err := a.ProfileService.CreateRawEducation(r.Context(), userID, services.CreateRawEducationParams{
		Label:       params.Label,
		School:      params.School,
		Description: params.Description,
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
