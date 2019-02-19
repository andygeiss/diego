package repositories

import (
	"errors"
	"github.com/andygeiss/diego/pkg/inference"
)

type defaultRepository struct {
	concepts []*inference.Concept
}

// NewDefaultRepository ...
func NewDefaultRepository(concepts []*inference.Concept) inference.Repository {
	return &defaultRepository{
		concepts: concepts,
	}
}

// FindQuestionsBySurvey ...
func (r *defaultRepository) FindConceptByName(name string) (*inference.Concept, error) {
	for _, concept := range r.concepts {
		if concept.Name == name {
			return concept, nil
		}
	}
	return nil, errors.New("concept was not found")
}
