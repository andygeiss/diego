package survey

import (
	"github.com/andygeiss/diego/pkg/explanation"
)

// Service ...
type Service interface {
	FindQuestionsBySurvey(name string) ([]*explanation.Question, error)
	GetResultsByFacts([]string) ([]string, error)
}
