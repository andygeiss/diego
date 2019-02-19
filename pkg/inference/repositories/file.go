package repositories

import (
	"encoding/json"
	"github.com/andygeiss/diego/pkg/inference"
	"io/ioutil"
)

type fileRepository struct {
	filename string
}

// NewFileRepository ...
func NewFileRepository(filename string) inference.Repository {
	return &fileRepository{
		filename: filename,
	}
}

// FindConceptByName ...
func (r *fileRepository) FindConceptByName(name string) (*inference.Concept, error) {

	var concepts []*inference.Concept

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
