package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/benKapl/cvmaker_api/internal/llm"
	"github.com/google/uuid"
)

type Offer struct {
	ID                      uuid.UUID `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Label                   string    `json:"label"`
	Organization            string    `json:"organization"`
	OrganizationDescription *string   `json:"organization_description,omitempty"`
	Missions                string    `json:"missions"`
	Stack                   *string   `json:"stack,omitempty"`
	ExpectedProfile         string    `json:"expected_profile"`
	Miscellaneous           *string   `json:"miscellaneous,omitempty"`
	UserID                  uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerOffersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type response struct {
		Offer Offer `json:"offer"`
	}

	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get userID from context", nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	llmOffer, err := llm.ParseOffer(params.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse offer", err)
		return
	}

	// IMPORTANT
	// Need to manage NULL STRINGS behavior !
	// Do that after creating LLM functions

	// Create optional fields with sql.NullString
	orgDescription := sql.NullString{String: "Une boite qui fait des trucs d'enculés", Valid: true}

	// Optional fields initialized as null
	if llmOffer.Stack == "" {
	}
	stack := sql.NullString{Valid: false}
	miscellaneous := sql.NullString{Valid: false}

	offer, err := cfg.db.CreateOffer(r.Context(), database.CreateOfferParams{
		Label:                   "FirstOffer",
		Organization:            "EnculéCorp",
		OrganizationDescription: orgDescription,
		Missions:                llmOffer.Missions,
		Stack:                   stack,
		ExpectedProfile:         llmOffer.ExpectedProfile,
		Miscellaneous:           miscellaneous,
		UserID:                  userID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create offer", err)
		return
	}

	// Convert database offer to response offer
	responseOffer := databaseOfferToOffer(offer)

	respondWithJSON(w, http.StatusCreated, response{
		Offer: responseOffer,
	})
}

func databaseOfferToOffer(dbOffer database.Offer) Offer {
	offer := Offer{
		ID:              dbOffer.ID,
		CreatedAt:       dbOffer.CreatedAt,
		UpdatedAt:       dbOffer.UpdatedAt,
		Label:           dbOffer.Label,
		Organization:    dbOffer.Organization,
		Missions:        dbOffer.Missions,
		ExpectedProfile: dbOffer.ExpectedProfile,
		UserID:          dbOffer.UserID,
	}

	if dbOffer.OrganizationDescription.Valid {
		offer.OrganizationDescription = &dbOffer.OrganizationDescription.String
	}

	if dbOffer.Stack.Valid {
		offer.Stack = &dbOffer.Stack.String
	}

	if dbOffer.Miscellaneous.Valid {
		offer.Miscellaneous = &dbOffer.Miscellaneous.String
	}

	return offer
}
