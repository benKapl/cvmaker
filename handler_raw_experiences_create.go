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

type Experience struct {
	Id           uuid.UUID    `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Title        string       `json:"title"`
	Organization string       `json:"organization"`
	Description  string       `json:"description"`
	Stacks       []string     `json:"stack,omitempty"`
	StartDate    time.Time    `json:"start_date"`
	EndDate      sql.NullTime `json:"end_date"`
	UserID       uuid.UUID    `json:"user_id"`
}

func (cfg *apiConfig) handlerRawExperiencesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title        string    `json:"title"`
		Organization string    `json:"organization"`
		Description  string    `json:"description"`
		Stacks       []string  `json:"stack,omitempty"`
		StartDate    time.Time `json:"start_date"`
		EndDate      time.Time `json:"end_date"`
		UserID       uuid.UUID `json:"user_id"`
	}

	type response struct {
		Success    bool `json:"success"`
		Experience Experience
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

	// Create dbExp in database
	dbExp, err := cfg.db.CreateRawExperience(r.Context(), database.CreateRawExperienceParams{
		Title:        strings.ToLower(params.Title),
		Organization: strings.ToLower(params.Organization),
		Description:  strings.ToLower(params.Description),
		StartDate:    params.StartDate,
		EndDate:      endDate,
		UserID:       userID,
	})

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "23505" { // Duplicate Key error
				respondWithError(w, http.StatusBadRequest, "User's raw experience already exists", err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create raw experience", err)
		return
	}

	// Manage stack retrieval and creation in pivot table
	// for stack in stacks:
	//		get stack in db
	// 		if not stack:
	//			create stack
	//      create expreience_stack
}
