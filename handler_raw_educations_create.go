package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (cfg *apiConfig) handlerRawEducationsCreate(w http.ResponseWriter, r *http.Request) {
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

	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Could not get userID, from Context", nil)
		return
	}

	var params parameters

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	var endDate sql.NullTime
	if params.EndDate.IsZero() {
		endDate = sql.NullTime{Valid: false}
	} else {
		endDate = sql.NullTime{Time: params.EndDate, Valid: true}
	}

	education, err := cfg.db.CreateRawEducation(r.Context(), database.CreateRawEducationParams{
		Label:       strings.ToLower(params.Label),
		School:      strings.ToLower(params.School),
		Description: strings.ToLower(params.Description),
		StartDate:   params.StartDate,
		EndDate:     endDate,
		UserID:      userID,
	})

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "23505" { // Duplicate Key error
				respondWithError(w, http.StatusBadRequest, "User's raw education already exists", err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create raw education", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Success: true,
		Education: Education{
			Id:          education.ID,
			CreatedAt:   education.CreatedAt,
			UpdatedAt:   education.UpdatedAt,
			Label:       education.Label,
			School:      education.School,
			Description: education.Description,
			StartDate:   education.StartDate,
			EndDate:     education.EndDate,
			UserID:      education.UserID,
		},
	})
}
