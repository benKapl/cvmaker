package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Offer struct {
	ID                      uuid.UUID `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Label                   string    `json:"label"`
	Organization            string    `json:"organization"`
	OrganizationDescription *string   `json:"organization_description"`
	Missions                string    `json:"missions"`
	Stack                   *string   `json:"stack"`
	ExpectedProfile         string    `json:"expected_profile"`
	Miscellaneous           *string   `json:"miscellaneous"`
	UserID                  uuid.UUID `json:"user_id"`
}

func handlerOffersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		body string
	}

	type response struct {
		Offer
	}
}
