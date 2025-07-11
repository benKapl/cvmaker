package services

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/mocks"
	"github.com/benKapl/cvmaker-api/internal/prompter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOfferService_CreateOffer(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()
	rawOffer := "some raw offer string"

	t.Run("successful offer creation", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		mockPrompter := new(mocks.MockPrompter)
		offerService := NewOfferService(mockDb, mockPrompter)

		parsedOffer := prompter.ParsedOffer{
			Title:           "Software Engineer",
			Organization:    "Google",
			Missions:        []string{"develop backend services"},
			Stack:           []string{"Go", "Kubernetes"},
			ExpectedProfile: []string{"3+ years experience"},
		}
		expectedDbOffer := database.Offer{
			ID:           uuid.New(),
			Title:        parsedOffer.Title,
			Organization: parsedOffer.Organization,
			Missions:     parsedOffer.Missions,
			Stack:        parsedOffer.Stack,
			UserID:       userID,
		}

		mockPrompter.On("ParseOffer", ctx, rawOffer).Return(parsedOffer, nil).Once()
		mockDb.On("CreateOffer", ctx, mock.AnythingOfType("database.CreateOfferParams")).Return(expectedDbOffer, nil).Once()

		createdOffer, err := offerService.CreateOffer(ctx, userID, rawOffer)

		assert.NoError(t, err)
		assert.Equal(t, expectedDbOffer, createdOffer)
		mockPrompter.AssertExpectations(t)
		mockDb.AssertExpectations(t)
	})

	t.Run("prompter returns error", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		mockPrompter := new(mocks.MockPrompter)
		offerService := NewOfferService(mockDb, mockPrompter)

		prompterErr := errors.New("failed to parse offer")
		mockPrompter.On("ParseOffer", ctx, rawOffer).Return(prompter.ParsedOffer{}, prompterErr).Once()

		_, err := offerService.CreateOffer(ctx, userID, rawOffer)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "couldn't parse offer")
		assert.ErrorIs(t, err, prompterErr)
		mockPrompter.AssertExpectations(t)
		mockDb.AssertNotCalled(t, "CreateOffer") // Ensure DB is not called if prompter fails
	})

	t.Run("database returns error", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		mockPrompter := new(mocks.MockPrompter)
		offerService := NewOfferService(mockDb, mockPrompter)

		parsedOffer := prompter.ParsedOffer{
			Title:           "Software Engineer",
			Organization:    "Google",
			Missions:        []string{"develop backend services"},
			Stack:           []string{"Go", "Kubernetes"},
			ExpectedProfile: []string{"3+ years experience"},
		}
		dbErr := errors.New("failed to create offer in db")

		mockPrompter.On("ParseOffer", ctx, rawOffer).Return(parsedOffer, nil).Once()
		mockDb.On("CreateOffer", ctx, mock.AnythingOfType("database.CreateOfferParams")).Return(database.Offer{}, dbErr).Once()

		_, err := offerService.CreateOffer(ctx, userID, rawOffer)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "couldn't create offer in db")
		assert.ErrorIs(t, err, dbErr)
		mockPrompter.AssertExpectations(t)
		mockDb.AssertExpectations(t)
	})

	t.Run("successful offer creation with optional fields", func(t *testing.T) {
		mockDb := new(mocks.MockQuerier)
		mockPrompter := new(mocks.MockPrompter)
		offerService := NewOfferService(mockDb, mockPrompter)

		orgDesc := "A leading tech company"
		parsedOffer := prompter.ParsedOffer{
			Title:                   "Software Engineer",
			Organization:            "Google",
			OrganizationDescription: &orgDesc,
			Missions:                []string{"develop backend services"},
			Stack:                   []string{"Go", "Kubernetes"},
			ExpectedProfile:         []string{"3+ years experience"},
			Miscellaneous:           []string{"Remote friendly"},
		}
		expectedDbOffer := database.Offer{
			ID:                      uuid.New(),
			Title:                   parsedOffer.Title,
			Organization:            parsedOffer.Organization,
			OrganizationDescription: sql.NullString{String: orgDesc, Valid: true},
			Missions:                parsedOffer.Missions,
			Stack:                   parsedOffer.Stack,
			ExpectedProfile:         parsedOffer.ExpectedProfile,
			Miscellaneous:           parsedOffer.Miscellaneous,
			UserID:                  userID,
		}

		mockPrompter.On("ParseOffer", ctx, rawOffer).Return(parsedOffer, nil).Once()
		mockDb.On("CreateOffer", ctx, mock.AnythingOfType("database.CreateOfferParams")).Return(expectedDbOffer, nil).Once()

		createdOffer, err := offerService.CreateOffer(ctx, userID, rawOffer)

		assert.NoError(t, err)
		assert.Equal(t, expectedDbOffer, createdOffer)
		mockPrompter.AssertExpectations(t)
		mockDb.AssertExpectations(t)
	})
}
