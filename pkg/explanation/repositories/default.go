package repositories

import (
	"errors"
	"github.com/andygeiss/diego/pkg/explanation"
)

type defaultRepository struct {
	surveys []*explanation.Survey
}

// NewDefaultRepository ...
func NewDefaultRepository(surveys []*explanation.Survey) explanation.Repository {
	return &defaultRepository{
		surveys: surveys,
	}
}

// FindQuestionsBySurvey ...
func (r *defaultRepository) FindQuestionsBySurvey(name string) ([]*explanation.Question, error) {
	var target *explanation.Survey
	for _, survey := range r.surveys {
		if survey.Name == name {
			target = survey
		}
	}
	if target == nil {
		return nil, errors.New("survey was not found")
	}
	return target.Questions, nil
}
