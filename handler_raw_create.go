package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
)

type Hobby struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerRawHobbiesCreate(w http.ResponseWriter, r *http.Request) {
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
		Label:  params.Label,
		UserID: userID,
	})
	if err != nil {
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
