package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Experience struct {
	Id           uuid.UUID    `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	Title        string       `json:"title"`
	Organization string       `json:"organization"`
	Description  string       `json:"description"`
	Stack        []string     `json:"stack,omitempty"`
	StartDate    time.Time    `json:"start_date"`
	EndDate      sql.NullTime `json:"end_date"`
	UserID       uuid.UUID    `json:"user_id"`
}

func (cfg *apiConfig) handlerRawExperiencesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Title        string       `json:"title"`
		Organization string       `json:"organization"`
		Description  string       `json:"description"`
		Stack        []string     `json:"stack,omitempty"`
		StartDate    time.Time    `json:"start_date"`
		EndDate      sql.NullTime `json:"end_date"`
		UserID       uuid.UUID    `json:"user_id"`
	}

	type response struct {
		Success    bool `json:"success"`
		Experience Experience
	}
}
