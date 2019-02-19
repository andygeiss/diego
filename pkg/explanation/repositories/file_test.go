package repositories_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/diego/pkg/explanation/repositories"
	"testing"
)

func TestFileRepository_FindQuestionsBySurvey_Should_Return_One_Question(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/explanation_1.json")
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(q), is.Equal(1))
	assert.That(t, q[0].Prompt, is.Equal("Go on Vacation?"))
}

func TestFileRepository_FindQuestionsBySurvey_Should_Return_Two_Questions(t *testing.T) {
	r := repositories.NewFileRepository("../../../testdata/explanation_2.json")
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(q), is.Equal(2))
	assert.That(t, q[1].Prompt, is.Equal("Go to Europe?"))
}
