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
	expRepo := explanation.NewDefaultRepository("../../app_explanation.json")
	infRepo := inference.NewDefaultRepository("../../app_inference.json")
	engine := inference.NewDefaultEngine("Schadensklasse bestimmen", infRepo)
	service := survey.NewDefaultService(expRepo, engine)
	qs, err := service.FindQuestionsBySurvey("Schadensklasse bestimmen")
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(qs), matchers.IsEqual(1))
	assert.That(t, qs[0].Prompt, matchers.IsEqual("Auswirkungen?"))
}

func TestDefaultService_GetResultsByFacts_Should_Return_One_Fact(t *testing.T) {
	expRepo := explanation.NewDefaultRepository("../../app_explanation.json")
	infRepo := inference.NewDefaultRepository("../../app_inference.json")
	engine := inference.NewDefaultEngine("Schadensklasse bestimmen", infRepo)
	service := survey.NewDefaultService(expRepo, engine)
	facts, err := service.GetResultsByFacts([]string{"#Q1 = 1"})
	assert.That(t, err, matchers.IsEqual(nil))
	assert.That(t, len(facts), matchers.IsEqual(1))
	assert.That(t, facts[0], matchers.IsEqual("keine nennenswerte Auswirkung auf die Geschäftstätigkeit"))
}
