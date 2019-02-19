package repositories_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/diego/pkg/explanation"
	"github.com/andygeiss/diego/pkg/explanation/repositories"
	"testing"
)

func arraySurveyWithOneQuestion() []*explanation.Survey {
	return []*explanation.Survey{
		&explanation.Survey{
			Name: "Survey",
			Questions: []*explanation.Question{
				&explanation.Question{
					Prompt: "Go on Vacation?",
					Options: []*explanation.Option{
						&explanation.Option{
							Name:  "Go to Europe",
							Value: "1",
						},
						&explanation.Option{
							Name:  "Visit Family",
							Value: "2",
						},
						&explanation.Option{
							Name:  "Go Camping Instate",
							Value: "3",
						},
					},
				},
			},
		},
	}
}

func arrangeSurveyWithTwoQuestions() []*explanation.Survey {
	return []*explanation.Survey{
		&explanation.Survey{
			Name: "Survey",
			Questions: []*explanation.Question{
				&explanation.Question{
					Prompt: "Go on Vacation?",
					Options: []*explanation.Option{
						&explanation.Option{
							Name:  "Go to Europe",
							Value: "1",
						},
						&explanation.Option{
							Name:  "Visit Family",
							Value: "2",
						},
						&explanation.Option{
							Name:  "Go Camping Instate",
							Value: "3",
						},
					},
				},
				&explanation.Question{
					Prompt: "Go to Europe?",
					Options: []*explanation.Option{
						&explanation.Option{
							Name:  "Visit Spain",
							Value: "1",
						},
						&explanation.Option{
							Name:  "Visit Germany",
							Value: "2",
						},
						&explanation.Option{
							Name:  "Visit Hungary",
							Value: "3",
						},
					},
				},
			},
		},
	}
}

func TestDefaultRepository_FindQuestionsBySurvey_Should_Return_One_Question(t *testing.T) {
	r := repositories.NewDefaultRepository(arraySurveyWithOneQuestion())
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(q), is.Equal(1))
	assert.That(t, q[0].Prompt, is.Equal("Go on Vacation?"))
}

func TestDefaultRepository_FindQuestionsBySurvey_Should_Return_Two_Questions(t *testing.T) {
	r := repositories.NewDefaultRepository(arrangeSurveyWithTwoQuestions())
	q, err := r.FindQuestionsBySurvey("Survey")
	assert.That(t, err, is.Equal(nil))
	assert.That(t, len(q), is.Equal(2))
	assert.That(t, q[1].Prompt, is.Equal("Go to Europe?"))
}
