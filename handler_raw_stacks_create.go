package main

import (
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
)

type Stack struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
	UserID    uuid.UUID `json:"user_id"`
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
