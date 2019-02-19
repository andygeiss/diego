package repositories_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/diego/pkg/inference/repositories"
	"testing"
)

func TestDefaultRepository_FindConceptByName_Should_Return_One_Concept(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/inference_1.json")
	c, err := r.FindConceptByName("one")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, c.Name, is.Equal("one"))
	assert.That(t, len(c.Facts), is.Equal(1))
	assert.That(t, len(c.Rules), is.Equal(1))
}
