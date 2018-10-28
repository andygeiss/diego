package main

import (
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/internal/inference"
	"github.com/andygeiss/diego/internal/survey"
	"github.com/andygeiss/diego/internal/survey/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	bind := os.Getenv("BIND")
	infRepoFile := os.Getenv("INF_REPO")
	expRepoFile := os.Getenv("EXP_REPO")
	surveyName := os.Getenv("SURVEY")

	log.Printf("INFO : Explanation Repository is   [%s]", expRepoFile)
	log.Printf("INFO : Inference Repository is     [%s]", infRepoFile)
	log.Printf("INFO : Survey name is              [%s]", surveyName)

	infRepo := inference.NewDefaultRepository(infRepoFile)
	engine := inference.NewDefaultEngine(surveyName, infRepo)
	expoRepo := explanation.NewDefaultRepository(expRepoFile)
	service := survey.NewDefaultService(expoRepo, engine)

	http.Handle("/questions", handlers.NewFindQuestionsBySurveyHandler(service))
	http.Handle("/results", handlers.NewGetResultsByFactsHandler(service))

	for {
		log.Printf("INFO : Starting Server ...")
		log.Printf("INFO : Listening at [%s] ...", bind)
		if err := http.ListenAndServe(bind, nil); err != nil {
			log.Printf("ERROR: ListenAndServe failed! [%s]", err.Error())
		}
	}
}
