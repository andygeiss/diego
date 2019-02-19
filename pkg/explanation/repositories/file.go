package repositories

import (
	"encoding/json"
	"errors"
	"github.com/andygeiss/diego/pkg/explanation"
	"io/ioutil"
)

type fileRepository struct {
	filename string
}

// NewFileRepository ...
func NewFileRepository(filename string) explanation.Repository {
	return &fileRepository{
		filename: filename,
	}
}

// FindQuestionsBySurvey ...
func (r *fileRepository) FindQuestionsBySurvey(name string) ([]*explanation.Question, error) {

	var surveys []*explanation.Survey

	data, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &surveys); err != nil {
		return nil, err
	}

	var target *explanation.Survey
	for _, survey := range surveys {
		if survey.Name == name {
			target = survey
		}
	}

	if target == nil {
		return nil, errors.New("survey was not found")
	}

	return target.Questions, nil
}
