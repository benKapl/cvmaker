package services

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProfileService struct {
	DB *database.Queries
}

func NewProfileService(db *database.Queries) *ProfileService {
	return &ProfileService{
		DB: db,
	}
}

type CreateRawEducationParams struct {
	Label       string
	School      string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}

func (s *ProfileService) CreateRawEducation(ctx context.Context, userID uuid.UUID, params CreateRawEducationParams) (database.RawEducation, error) {
	var endDate sql.NullTime
	if !params.EndDate.IsZero() {
		endDate = sql.NullTime{Time: params.EndDate, Valid: true}
	}

	rawEducation, err := s.DB.CreateRawEducation(ctx, database.CreateRawEducationParams{
		Label:       strings.ToLower(params.Label),
		School:      strings.ToLower(params.School),
		Description: strings.ToLower(params.Description),
		StartDate:   params.StartDate,
		EndDate:     endDate,
		UserID:      userID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.RawEducation{}, ErrDuplicateKey
			}
		}
		return database.RawEducation{}, fmt.Errorf("couldn't create rawEducation: %w", err)
	}

	return rawEducation, nil
}
