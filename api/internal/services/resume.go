package services

import (
	"github.com/benKapl/cvmaker-api/internal/database"
	"github.com/benKapl/cvmaker-api/internal/llm"
)

type ResumeService struct {
	DB        *database.Queries
	LLMClient llm.LLMClient
}

func NewResumeService(db *database.Queries, llmClient llm.LLMClient) *ResumeService {
	return &ResumeService{
		DB:        db,
		LLMClient: llmClient,
	}
}
