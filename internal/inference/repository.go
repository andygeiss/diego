package inference

import (
	"encoding/json"
	"io/ioutil"
)

// Concept ...
type Concept struct {
	Name string
	Facts []string
	Rules []*Rule
}

// Repository ...
type Repository interface {
	FindConceptByName(name string) (*Concept, error)
}

// Rule ...
type Rule struct {
	Conditions  []string
	Conclusions []string
}

type defaultRepository struct {
	filename string
}

// NewDefaultRepository ...
func NewDefaultRepository(filename string) Repository {
	return &defaultRepository{
		filename: filename,
	}
}

// FindConceptByName ...
func (r *defaultRepository) FindConceptByName(name string) (*Concept, error) {

	var concepts []*Concept

	data, err := ioutil.ReadFile(r.filename)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &concepts); err != nil {
		return nil, err
	}

	for _, concept := range concepts {
		if concept.Name == name {
			return concept, nil
		}
	}

	return nil, nil
}
