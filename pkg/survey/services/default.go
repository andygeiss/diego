package services

import (
	"github.com/andygeiss/diego/pkg/explanation"
	"github.com/andygeiss/diego/pkg/inference"
	"github.com/andygeiss/diego/pkg/survey"
	"strings"
)

type defaultService struct {
	repo   explanation.Repository
	engine inference.Engine
}

// NewDefaultService ...
func NewDefaultService(repo explanation.Repository, engine inference.Engine) survey.Service {
	return &defaultService{
		repo:   repo,
		engine: engine,
	}
}

// FindSurveyByName ...
func (s *defaultService) FindQuestionsBySurvey(name string) ([]*explanation.Question, error) {
	return s.repo.FindQuestionsBySurvey(name)
}

// GetResultsByFacts ...
func (s *defaultService) GetResultsByFacts(facts []string) ([]string, error) {

	result, err := s.engine.Run(facts)
	if err != nil {
		return nil, err
	}

	filtered := make([]string, 0)
	for _, fact := range result {
		if !isNotAQuestion(fact) {
			filtered = append(filtered, fact)
		}
	}

	return filtered, nil
}

func isNotAQuestion(fact string) bool {
	return strings.HasPrefix(fact, "#Q")
}
