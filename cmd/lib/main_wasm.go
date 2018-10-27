package main

import (
	"fmt"
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/internal/survey/handlers"
	"github.com/andygeiss/diego/pkg/wasm"
	"log"
	"syscall/js"
)

const (
	questionsURL = "http://127.0.0.1:3000/questions"
	resultsURL   = "http://127.0.0.1:3000/results"
	surveyName   = "Schadensklasse bestimmen"
)

var selection = make(map[string]string, 0)

func main() {
	c := make(chan struct{}, 0)
	log.Printf("INFO : App is running ...")
	prepareSurvey(surveyName)
	prepareEvaluation()
	showContents()
	<-c
}

func prepareSurvey(name string) {
	btnLoad := wasm.GetById("btnLoad")
	// Onclick
	btnLoad.Call("addEventListener", "click", js.NewCallback(func(v []js.Value) {
		go func() {
			req := &handlers.FindQuestionsBySurveyRequest{Name: name}
			res, err := handlers.InvokeFindQuestionsBySurvey(questionsURL, req)
			if err != nil {
				updateErrors(err)
				return
			}
			updateSurvey(res.Questions)
		}()
	}))
	// Make visible => ready
	btnLoad.Set("style", "display: block")
}

func prepareEvaluation() {
	btnEvaluate := wasm.GetById("btnEvaluate")
	// Onclick
	btnEvaluate.Call("addEventListener", "click", js.NewCallback(func(v []js.Value) {
		facts := make([]string, 0)
		for _, fact := range selection {
			facts = append(facts, fact)
		}
		go func() {
			req := &handlers.GetResultsByFactsRequest{Facts: facts}
			res, err := handlers.InvokeGetResultsByFacts(resultsURL, req)
			if err != nil {
				updateErrors(err)
				return
			}
			updateEvaluation(res.Results)
		}()
	}))
}

func showContents() {
	wasm.GetById("loader").Set("style", "display: none")
	wasm.GetById("content").Set("style", "display: flex")
}

func updateSurvey(questions []*explanation.Question) {
	parent := wasm.GetById("survey")
	// Clear previous inputs
	parent.Set("innerText", "")
	wasm.GetById("errors").Set("style", "display: none")
	wasm.GetById("results").Set("innerText", "")
	// Show button for evaluation
	wasm.GetById("btnEvaluate").Set("style", "display: block")
	// Show questions
	for i, question := range questions {
		qid := fmt.Sprintf("q%d", i+1)
		hidden := wasm.Create("input")
		hidden.Set("id", qid)
		hidden.Set("name", qid)
		hidden.Set("type", "hidden")
		prompt := wasm.Create("span")
		prompt.Set("innerText", question.Prompt)
		questionDiv := wasm.Create("div")
		questionDiv.Set("classList", "question")
		questionDiv.Call("appendChild", hidden)
		questionDiv.Call("appendChild", prompt)
		for j, option := range question.Options {
			oid := fmt.Sprintf("o%d", j+1)
			label := wasm.Create("span")
			label.Set("innerText", option.Name)
			optionDiv := wasm.Create("div")
			optionDiv.Set("classList", "option")
			optionDiv.Set("id", qid+oid)
			optionDiv.Set("value", option.Value)
			optionDiv.Call("appendChild", label)
			questionDiv.Call("appendChild", optionDiv)
		}
		parent.Call("appendChild", questionDiv)
	}
	// Add onclick events to options
	for i, question := range questions {
		qid := fmt.Sprintf("q%d", i+1)
		for j, _ := range question.Options {
			oid := fmt.Sprintf("o%d", j+1)
			optionDiv := wasm.GetById(qid + oid)
			optionDiv.Call("addEventListener", "click", js.NewCallback(func(v []js.Value) {
				selection[qid] = optionDiv.Get("value").String()
				// Clear previous selections
				for j, _ := range question.Options {
					wasm.GetById(qid+fmt.Sprintf("o%d", j+1)).Set("classList", "option")
				}
				// Highlight selected option
				wasm.GetById(qid+oid).Set("classList", "option selected")
			}))
		}
	}
}

func updateEvaluation(results []string) {
	parent := wasm.GetById("results")
	// Clear previous results
	parent.Set("innerText", "")
	wasm.GetById("errors").Set("style", "display: none")
	wasm.GetById("survey").Set("innerText", "")
	// Hide evaluation button
	wasm.GetById("btnEvaluate").Set("style", "display: none")
	// Show facts/results
	for _, fact := range results {
		factDiv := wasm.Create("div")
		factDiv.Set("classList", "fact")
		factDiv.Set("innerHTML", fact)
		parent.Call("appendChild", factDiv)
	}
}

func updateErrors(err error) {
	errDiv := wasm.GetById("errors")
	errDiv.Set("innerText", err.Error())
	errDiv.Set("style", "display: block")
}
