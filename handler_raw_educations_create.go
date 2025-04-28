package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Educations struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Label       string    `json:"label"`
	School      string    `json:"school"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UserID      uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerRawEducationsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label string `json:"label"`
	}

	type response struct {
		Success bool
		Hobby   Hobby
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

	hobby, err := cfg.db.CreateRawHobby(r.Context(), database.CreateRawHobbyParams{
		Label:  strings.ToLower(params.Label),
		UserID: userID,
	})

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "23505" { // Duplicate Key error
				respondWithError(w, http.StatusBadRequest, "User's raw hobby already exists", err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create raw hobby", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Success: true,
		Hobby: Hobby{
			Id:        hobby.ID,
			CreatedAt: hobby.CreatedAt,
			UpdatedAt: hobby.UpdatedAt,
			Label:     hobby.Label,
			UserID:    hobby.UserID,
		},
	})
}
