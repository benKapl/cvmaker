package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Stack struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerRawStacksCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Label string `json:"label"`
	}

	type response struct {
		Success bool `json:"success"`
		Stack   Stack
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

	dbStack, err := cfg.db.CreateRawStack(r.Context(), database.CreateRawStackParams{
		Label:  strings.ToLower(params.Label),
		UserID: userID,
	})

	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok {
			if pqErr.Code == "23505" { // Duplicate Key error
				respondWithError(w, http.StatusBadRequest, "User's raw stack already exists", err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create raw stack", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Success: true,
		Stack:   dbRawStackToStack(dbStack),
	})
}

func dbRawStackToStack(dbRawStack database.RawStack) Stack {
	return Stack{
		Id:        dbRawStack.ID,
		CreatedAt: dbRawStack.CreatedAt,
		UpdatedAt: dbRawStack.UpdatedAt,
		Label:     dbRawStack.Label,
		UserID:    dbRawStack.UserID,
	}
}
