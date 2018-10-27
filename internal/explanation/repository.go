package explanation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Option ...
type Option struct {
	Name string
	Value interface{}
}

// Question ...
type Question struct {
	Prompt  string
	Options []*Option
}

// Repository ...
type Repository interface {
	FindQuestionsBySurvey(name string) ([]*Question, error)
}

// Survey ...
type Survey struct {
	Name string
	Questions []*Question
}

type defaultRepository struct {
	filename string
}

// NewDefaultRepository ...
func NewDefaultRepository(filename string) Repository {
	return &defaultRepository{
		filename:filename,
	}
}

// FindQuestionsBySurvey ...
func (r *defaultRepository) FindQuestionsBySurvey(name string) ([]*Question, error) {

	var surveys []*Survey

	data, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &surveys); err != nil {
		return nil, err
	}

	var target *Survey
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
