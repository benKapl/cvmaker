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
	DB database.Storer
}

func NewProfileService(db database.Storer) *ProfileService {
	return &ProfileService{
		DB: db,
	}
}

type CreateRawEducationParams struct {
	Label       string
	School      string
	Description string
	StartDate   *time.Time
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
		StartDate:   *params.StartDate,
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

type CreateRawStackParams struct {
	Label string
}

func (s *ProfileService) CreateRawStack(ctx context.Context, userID uuid.UUID, params CreateRawStackParams) (database.RawStack, error) {
	rawStack, err := s.DB.CreateRawStack(ctx, database.CreateRawStackParams{
		Label:  strings.ToLower(params.Label),
		UserID: userID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.RawStack{}, ErrDuplicateKey
			}
		}
		return database.RawStack{}, fmt.Errorf("couldn't create rawStack: %w", err)
	}

	return rawStack, nil
}

type CreateRawHobbyParams struct {
	Label string
}

func (s *ProfileService) CreateRawHobby(ctx context.Context, userID uuid.UUID, params CreateRawHobbyParams) (database.RawHobby, error) {
	rawHobby, err := s.DB.CreateRawHobby(ctx, database.CreateRawHobbyParams{
		Label:  strings.ToLower(params.Label),
		UserID: userID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.RawHobby{}, ErrDuplicateKey
			}
		}
		return database.RawHobby{}, fmt.Errorf("couldn't create rawHobby: %w", err)
	}

	return rawHobby, nil
}

type CreateRawProjectParams struct {
	Label       string
	Description string
	Stacks      []string
	StartDate   *time.Time
	EndDate     time.Time
}

func (s *ProfileService) CreateRawProject(ctx context.Context, userID uuid.UUID, params CreateRawProjectParams) (database.RawProject, []database.RawStack, error) {
	if params.StartDate == nil {
		return database.RawProject{}, nil, fmt.Errorf("start_date is required")
	}

	var endDate sql.NullTime
	if !params.EndDate.IsZero() {
		endDate = sql.NullTime{Time: params.EndDate, Valid: true}
	}

	rawProject, err := s.DB.CreateRawProject(ctx, database.CreateRawProjectParams{
		Label:       strings.ToLower(params.Label),
		Description: strings.ToLower(params.Description),
		StartDate:   *params.StartDate,
		EndDate:     endDate,
		UserID:      userID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.RawProject{}, nil, ErrDuplicateKey
			}
		}
		return database.RawProject{}, nil, fmt.Errorf("couldn't create rawProject: %w", err)
	}

	var stacks []database.RawStack
	if len(params.Stacks) > 0 {
		for _, stackLabel := range params.Stacks {
			dbStack, err := s.DB.GetRawStackByLabel(ctx, database.GetRawStackByLabelParams{
				Label:  strings.ToLower(stackLabel),
				UserID: userID,
			})
			if err != nil && err != sql.ErrNoRows {
				return database.RawProject{}, nil, fmt.Errorf("couldn't get raw stack: %w", err)
			}

			if err == sql.ErrNoRows {
				dbStack, err = s.DB.CreateRawStack(ctx, database.CreateRawStackParams{
					Label:  strings.ToLower(stackLabel),
					UserID: userID,
				})
				if err != nil {
					return database.RawProject{}, nil, fmt.Errorf("couldn't create raw stack: %w", err)
				}
			}

			_, err = s.DB.CreateRawProjectStack(ctx, database.CreateRawProjectStackParams{
				ProjectID: rawProject.ID,
				StackID:   dbStack.ID,
			})
			if err != nil {
				return database.RawProject{}, nil, fmt.Errorf("couldn't create raw project_stack: %w", err)
			}

			stacks = append(stacks, dbStack)
		}
	}

	return rawProject, stacks, nil
}

type CreateRawExperienceParams struct {
	Title        string
	Organization string
	Description  string
	Stacks       []string
	StartDate    *time.Time
	EndDate      time.Time
}

func (s *ProfileService) CreateRawExperience(ctx context.Context, userID uuid.UUID, params CreateRawExperienceParams) (database.RawExperience, []database.RawStack, error) {
	if params.StartDate == nil {
		return database.RawExperience{}, nil, fmt.Errorf("start_date is required")
	}

	var endDate sql.NullTime
	if !params.EndDate.IsZero() {
		endDate = sql.NullTime{Time: params.EndDate, Valid: true}
	}

	rawExperience, err := s.DB.CreateRawExperience(ctx, database.CreateRawExperienceParams{
		Title:        strings.ToLower(params.Title),
		Organization: strings.ToLower(params.Organization),
		Description:  strings.ToLower(params.Description),
		StartDate:    *params.StartDate,
		EndDate:      endDate,
		UserID:       userID,
	})

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return database.RawExperience{}, nil, ErrDuplicateKey
			}
		}
		return database.RawExperience{}, nil, fmt.Errorf("couldn't create rawExperience: %w", err)
	}

	var stacks []database.RawStack
	if len(params.Stacks) > 0 {
		for _, stackLabel := range params.Stacks {
			dbStack, err := s.DB.GetRawStackByLabel(ctx, database.GetRawStackByLabelParams{
				Label:  strings.ToLower(stackLabel),
				UserID: userID,
			})
			if err != nil && err != sql.ErrNoRows {
				return database.RawExperience{}, nil, fmt.Errorf("couldn't get raw stack: %w", err)
			}

			if err == sql.ErrNoRows {
				dbStack, err = s.DB.CreateRawStack(ctx, database.CreateRawStackParams{
					Label:  strings.ToLower(stackLabel),
					UserID: userID,
				})
				if err != nil {
					return database.RawExperience{}, nil, fmt.Errorf("couldn't create raw stack: %w", err)
				}
			}

			_, err = s.DB.CreateRawExperienceStack(ctx, database.CreateRawExperienceStackParams{
				ExperienceID: rawExperience.ID,
				StackID:      dbStack.ID,
			})
			if err != nil {
				return database.RawExperience{}, nil, fmt.Errorf("couldn't create raw experience_stack: %w", err)
			}

			stacks = append(stacks, dbStack)
		}
	}

	return rawExperience, stacks, nil
}
