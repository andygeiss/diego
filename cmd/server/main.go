package main

import (
	"flag"
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/internal/inference"
	"github.com/andygeiss/diego/internal/survey"
	"github.com/andygeiss/diego/internal/survey/handlers"
	"log"
	"net/http"
)

func main() {

	address := ":3000"

	infRepoFile := flag.String("inf", "", "JSON-File of Inference Repository")
	expRepoFile := flag.String("exp", "", "JSON-File of Explanation Repository")
	surveyName := flag.String("survey", "", "Name of the Survey")
	flag.Parse()

	infRepo := inference.NewDefaultRepository(*infRepoFile)
	engine := inference.NewDefaultEngine(*surveyName, infRepo)
	expoRepo := explanation.NewDefaultRepository(*expRepoFile)
	service := survey.NewDefaultService(expoRepo, engine)

	http.Handle("/questions", handlers.NewFindQuestionsBySurveyHandler(service))
	http.Handle("/results", handlers.NewGetResultsByFactsHandler(service))

	for {
		log.Printf("INFO : Starting Server ...")
		log.Printf("INFO : Listening at [%s] ...", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Printf("ERROR: ListenAndServe failed! [%s]", err.Error())
		}
	}
}
