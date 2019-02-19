package services_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	expRepos "github.com/andygeiss/diego/pkg/explanation/repositories"
	"github.com/andygeiss/diego/pkg/inference/engines"
	infRepos "github.com/andygeiss/diego/pkg/inference/repositories"
	"github.com/andygeiss/diego/pkg/survey/services"
	"testing"
)

func TestDefaultService_FindSurveyByName_Should_Return_One_Survey(t *testing.T) {
	expRepo := expRepos.NewFileRepository("../../../testdata/explanation.json")
	infRepo := infRepos.NewFileRepository("../../../testdata/inference.json")
	engine := engines.NewDefaultEngine("SURVEY NAME", infRepo)
	service := services.NewDefaultService(expRepo, engine)
	qs, err := service.FindQuestionsBySurvey("SURVEY NAME")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(qs), is.Equal(2))
	assert.That(t, qs[0].Prompt, is.Equal("Question 1 ?"))
}

func TestDefaultService_GetResultsByFacts_Should_Return_One_Fact(t *testing.T) {
	expRepo := expRepos.NewFileRepository("../../../testdata/explanation.json")
	infRepo := infRepos.NewFileRepository("../../../testdata/inference.json")
	engine := engines.NewDefaultEngine("SURVEY NAME", infRepo)
	service := services.NewDefaultService(expRepo, engine)
	facts, err := service.GetResultsByFacts([]string{"#Q1 = 1"})
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(facts), is.Equal(2))
	assert.That(t, facts[0], is.Equal("Question 1 answered"))
}
