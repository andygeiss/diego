package repositories_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/diego/pkg/inference"
	"github.com/andygeiss/diego/pkg/inference/repositories"
	"testing"
)

func arrangeOneConceptWithOneRule() []*inference.Concept {
	return []*inference.Concept{
		&inference.Concept{
			Name:  "one",
			Facts: []string{"A"},
			Rules: []*inference.Rule{
				&inference.Rule{
					Conditions:  []string{"A"},
					Conclusions: []string{"X"},
				},
			},
		},
	}
}

func arrangeOneConceptWithTwoRules() []*inference.Concept {
	return []*inference.Concept{
		&inference.Concept{
			Name:  "two",
			Facts: []string{"A", "B"},
			Rules: []*inference.Rule{
				&inference.Rule{
					Conditions:  []string{"A"},
					Conclusions: []string{"X"},
				},
				&inference.Rule{
					Conditions:  []string{"B"},
					Conclusions: []string{"Y"},
				},
			},
		},
	}
}

func TestDefaultRepository_FindConceptByName_Should_Return_One_Concept_With_One_Rule(t *testing.T) {
	r := repositories.NewDefaultRepository(arrangeOneConceptWithOneRule())
	c, err := r.FindConceptByName("one")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, c.Name, is.Equal("one"))
	assert.That(t, len(c.Facts), is.Equal(1))
	assert.That(t, len(c.Rules), is.Equal(1))
}

func TestDefaultRepository_FindConceptByName_Should_Return_One_Concept_With_Two_Rules(t *testing.T) {
	r := repositories.NewDefaultRepository(arrangeOneConceptWithTwoRules())
	c, err := r.FindConceptByName("two")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, c.Name, is.Equal("two"))
	assert.That(t, len(c.Facts), is.Equal(2))
	assert.That(t, len(c.Rules), is.Equal(2))
}
