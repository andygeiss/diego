package survey_test

import (
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/internal/inference"
	"github.com/andygeiss/diego/internal/survey"
	"github.com/andygeiss/diego/pkg/assert"
	"github.com/andygeiss/diego/pkg/assert/matchers"
	"testing"
)

func TestDefaultService_FindSurveyByName_Should_Return_One_Survey(t *testing.T) {
	expRepo := explanation.NewDefaultRepository("../../testdata/explanation.json")
	infRepo := inference.NewDefaultRepository("../../testdata/inference.json")
	engine := inference.NewDefaultEngine("SURVEY NAME", infRepo)
	service := survey.NewDefaultService(expRepo, engine)
	qs, err := service.FindQuestionsBySurvey("SURVEY NAME")
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(qs), matchers.IsEqual(2))
	assert.That(t, qs[0].Prompt, matchers.IsEqual("Question 1 ?"))
}

func TestDefaultService_GetResultsByFacts_Should_Return_One_Fact(t *testing.T) {
	expRepo := explanation.NewDefaultRepository("../../testdata/explanation.json")
	infRepo := inference.NewDefaultRepository("../../testdata/inference.json")
	engine := inference.NewDefaultEngine("SURVEY NAME", infRepo)
	service := survey.NewDefaultService(expRepo, engine)
	facts, err := service.GetResultsByFacts([]string{"#Q1 = 1"})
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(2))
	assert.That(t, facts[0], matchers.IsEqual("Question 1 answered"))
}
