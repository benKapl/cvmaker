package services

import (
	"context"

	"github.com/benKapl/cvmaker-api/internal/database"
)

type AdminService struct {
	DB       *database.Queries
	Platform string
}

func NewAdminService(db *database.Queries, platform string) *AdminService {
	return &AdminService{
		DB:       db,
		Platform: platform,
	}
}

func (s *AdminService) ResetDatabase(ctx context.Context) error {
	err := s.DB.DeleteUsers(ctx)
	if err != nil {
		return err
	}

	return nil
}
