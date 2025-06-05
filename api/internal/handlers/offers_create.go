package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker_api/internal/database"
	"github.com/google/uuid"
)

type Offer struct {
	ID                      uuid.UUID `json:"id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	Title                   string    `json:"title"`
	Organization            string    `json:"organization"`
	OrganizationDescription *string   `json:"organization_description,omitempty"` // Optional field, pointer allows nil
	Missions                []string  `json:"missions"`                           // NOT NULL in DB, but can be an empty array
	Stack                   []string  `json:"stack,omitempty"`                    // Optional array field
	ExpectedProfile         []string  `json:"expected_profile"`                   // NOT NULL in DB, but can be an empty array
	Miscellaneous           []string  `json:"miscellaneous,omitempty"`            // Optional array field
	UserID                  uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerOffersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type response struct {
		Success bool  `json:"success"`
		Offer   Offer `json:"offer"`
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

	// Call the LLM to parse the offer
	llmOffer, err := cfg.llmClient.ParseOffer(params.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse offer", err)
		return
	}

	dbParams := database.CreateOfferParams{
		Title:           llmOffer.Title,
		Organization:    llmOffer.Organization,
		Missions:        llmOffer.Missions,
		Stack:           llmOffer.Stack,
		ExpectedProfile: llmOffer.ExpectedProfile,
		Miscellaneous:   llmOffer.Miscellaneous,
		UserID:          userID,
	}
	// Map optional string field (OrganizationDescription) from *string to sql.NullString
	if llmOffer.OrganizationDescription != nil {
		dbParams.OrganizationDescription = sql.NullString{String: *llmOffer.OrganizationDescription, Valid: true}
	} else {
		dbParams.OrganizationDescription = sql.NullString{Valid: false}
	}

	// Create offer in database
	dbOffer, err := cfg.db.CreateOffer(r.Context(), dbParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create offer", err)
		return
	}

	offer := dbOfferToOffer(dbOffer)

	respondWithJSON(w, http.StatusCreated, response{
		Success: true,
		Offer:   offer,
	})
}

func dbOfferToOffer(dbOffer database.Offer) Offer {
	offer := Offer{
		ID:              dbOffer.ID,
		CreatedAt:       dbOffer.CreatedAt,
		UpdatedAt:       dbOffer.UpdatedAt,
		Title:           dbOffer.Title,
		Organization:    dbOffer.Organization,
		Missions:        dbOffer.Missions,
		ExpectedProfile: dbOffer.ExpectedProfile,
		UserID:          dbOffer.UserID,
	}
	// Map optional string description from sql.NullString to *string for API response
	if dbOffer.OrganizationDescription.Valid {
		offer.OrganizationDescription = &dbOffer.OrganizationDescription.String
	}
	offer.Stack = dbOffer.Stack
	offer.Miscellaneous = dbOffer.Miscellaneous

	return offer
}
