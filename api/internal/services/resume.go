package services

import (
	"log"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/prompter"
)

type ResumeService struct {
	DB       database.Querier
	Prompter prompter.Prompter
}

func NewResumeService(db database.Querier, p prompter.Prompter) *ResumeService {
	return &ResumeService{
		DB:       db,
		Prompter: p,
	}
}

func (s *ResumeService) CreateResume() {
	log.Println("WESH")
}
