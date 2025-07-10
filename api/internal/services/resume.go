package services

import (
	"log"

	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/llm"
)

type ResumeService struct {
	DB        database.Storer
	LLMClient llm.LLMClient
}

func NewResumeService(db database.Storer, llmClient llm.LLMClient) *ResumeService {
	return &ResumeService{
		DB:        db,
		LLMClient: llmClient,
	}
}

func (s *ResumeService) CreateResume() {
	log.Println("WESH")
}
