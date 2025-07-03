package services

import "github.com/benKapl/cvmaker-api/internal/database"

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

func (s *AdminService) ResetDatabase() string {
	return s.Platform
}
