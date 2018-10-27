package inference_test

import (
	"github.com/andygeiss/diego/internal/inference"
	"github.com/andygeiss/diego/pkg/assert"
	"github.com/andygeiss/diego/pkg/assert/matchers"
	"testing"
)

func TestDefaultRepository_FindConceptByName_Should_Return_One_Concept(t *testing.T) {
	r := inference.NewDefaultRepository("../../testdata/inference_1.json")
	c, err := r.FindConceptByName("one")
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, c.Name, matchers.IsEqual("one"))
	assert.That(t, len(c.Facts), matchers.IsEqual(1))
	assert.That(t, len(c.Rules), matchers.IsEqual(1))
}
