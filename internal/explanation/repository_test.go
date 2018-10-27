package explanation_test

import (
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/pkg/assert"
	"github.com/andygeiss/diego/pkg/assert/matchers"
	"testing"
)

func TestDefaultRepository_FindQuestionsBySurvey_Should_Return_One_Question(t *testing.T) {
	r := explanation.NewDefaultRepository("../../testdata/explanation_1.json")
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(q), matchers.IsEqual(1))
	assert.That(t, q[0].Prompt, matchers.IsEqual("Go on Vacation?"))
}

func TestDefaultRepository_FindQuestionsBySurvey_Should_Return_Two_Questions(t *testing.T) {
	r := explanation.NewDefaultRepository("../../testdata/explanation_2.json")
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(q), matchers.IsEqual(2))
	assert.That(t, q[1].Prompt, matchers.IsEqual("Go to Europe?"))
}
