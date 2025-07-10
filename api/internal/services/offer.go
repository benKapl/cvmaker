package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/prompter"
	"github.com/google/uuid"
)

type OfferService struct {
	DB       database.Querier
	Prompter prompter.Prompter
}

func NewOfferService(db database.Querier, p prompter.Prompter) *OfferService {
	return &OfferService{
		DB:       db,
		Prompter: p,
	}
}

func (s *OfferService) CreateOffer(ctx context.Context, userID uuid.UUID, rawOffer string) (database.Offer, error) {
	parsedOffer, err := s.Prompter.ParseOffer(ctx, rawOffer)
	if err != nil {
		return database.Offer{}, fmt.Errorf("couldn't parse offer: %w", err)
	}

	dbParams := database.CreateOfferParams{
		Title:           parsedOffer.Title,
		Organization:    parsedOffer.Organization,
		Missions:        parsedOffer.Missions,
		Stack:           parsedOffer.Stack,
		ExpectedProfile: parsedOffer.ExpectedProfile,
		Miscellaneous:   parsedOffer.Miscellaneous,
		UserID:          userID,
	}
	// Map optional string field (OrganizationDescription) from *string to sql.NullString
	if parsedOffer.OrganizationDescription != nil {
		dbParams.OrganizationDescription = sql.NullString{String: *parsedOffer.OrganizationDescription, Valid: true}
	} else {
		dbParams.OrganizationDescription = sql.NullString{Valid: false}
	}

	// Create offer in database
	offer, err := s.DB.CreateOffer(ctx, dbParams)
	if err != nil {
		return database.Offer{}, fmt.Errorf("couldn't create offer in db: %w", err)
	}

	return offer, nil
}