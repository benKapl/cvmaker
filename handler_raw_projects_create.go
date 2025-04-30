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

type Project struct {
	Id          uuid.UUID    `json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Label       string       `json:"label"`
	Description string       `json:"description"`
	Stacks      []Stack      `json:"stacks,omitempty"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     sql.NullTime `json:"end_date"`
	UserID      uuid.UUID    `json:"user_id"`
}

func (cfg *apiConfig) handlerRawProjectsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label       string    `json:"label"`
		Description string    `json:"description"`
		Stacks      []string  `json:"stacks,omitempty"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		UserID      uuid.UUID `json:"user_id"`
	}

	type response struct {
		Success bool `json:"success"`
		Project Project
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
	dbProject, err := cfg.db.CreateRawProject(r.Context(), database.CreateRawProjectParams{
		Label:       strings.ToLower(params.Label),
		Description: strings.ToLower(params.Description),
		StartDate:   params.StartDate,
		EndDate:     endDate,
		UserID:      userID,
	})

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "23505" { // Duplicate Key error
				respondWithError(w, http.StatusBadRequest, "User's raw project already exists", err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create raw project", err)
		return
	}

	// Manage stack retrieval and creation in pivot table
	var stacks []Stack

	if len(params.Stacks) > 0 {
		for _, stack := range params.Stacks {
			// Get stack in db
			dbStack, err := cfg.db.GetRawStackByLabel(r.Context(), database.GetRawStackByLabelParams{
				Label:  strings.ToLower(stack),
				UserID: userID,
			})
			if err != sql.ErrNoRows && err != nil {
				respondWithError(w, http.StatusInternalServerError, "Could not get raw stack but should have been able to", err)
				return
			}

			// if stack does not exist, create it
			if err == sql.ErrNoRows {
				dbStack, err = cfg.db.CreateRawStack(r.Context(), database.CreateRawStackParams{
					Label:  strings.ToLower(stack),
					UserID: userID,
				})
				if err != nil {
					respondWithError(w, http.StatusInternalServerError, "Could not create raw stack", err)
					return
				}
			}

			// Link stack to project
			_, err = cfg.db.CreateRawProjectStack(r.Context(), database.CreateRawProjectStackParams{
				ProjectID: dbProject.ID,
				StackID:   dbStack.ID,
			})
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "Could not create raw experience_stack", err)
				return
			}

			// If all went well, convert dbStack to stack and append it to stack list
			stacks = append(stacks, dbRawStackToStack(dbStack))
		}
	}

	respondWithJSON(w, http.StatusCreated, response{
		Success: true,
		Project: Project{
			Id:          dbProject.ID,
			CreatedAt:   dbProject.CreatedAt,
			UpdatedAt:   dbProject.UpdatedAt,
			Label:       dbProject.Label,
			Description: dbProject.Description,
			Stacks:      stacks,
			StartDate:   dbProject.StartDate,
			EndDate:     dbProject.EndDate,
			UserID:      dbProject.UserID,
		},
	})
}
