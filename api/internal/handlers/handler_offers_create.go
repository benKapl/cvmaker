package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/respond"
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

func (a *API) handlerOffersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Offer string `json:"offer"`
	}

	type response struct {
		Success bool  `json:"success"`
		Offer   Offer `json:"offer"`
	}

	userID, ok := r.Context().Value(userIDKey).(uuid.UUID)
	if !ok {
		respond.WithError(w, http.StatusUnauthorized, "Couldn't get userID from context", nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	dbOffer, err := a.OfferService.CreateOffer(r.Context(), userID, params.Offer)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Couldn't create offer", err)
		return
	}

	offer := dbOfferToOffer(dbOffer)

	respond.WithJSON(w, http.StatusCreated, response{
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
