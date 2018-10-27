package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/andygeiss/diego/internal/explanation"
	"github.com/andygeiss/diego/internal/survey"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type findQuestionsBySurveyHandler struct {
	service survey.Service
}

// FindQuestionsBySurveyRequest ...
type FindQuestionsBySurveyRequest struct {
	Name string `json:"name"`
}

// FindQuestionsBySurveyResponse ...
type FindQuestionsBySurveyResponse struct {
	Questions []*explanation.Question `json:"questions"`
}

// NewFindQuestionsBySurveyHandler ...
func NewFindQuestionsBySurveyHandler(service survey.Service) http.Handler {
	return &findQuestionsBySurveyHandler{service: service}
}

// ServeHTTP ...
func (h *findQuestionsBySurveyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Allows calls from WebAssembly to different ports.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: ReadAll HTML body failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	in, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		log.Printf("ERROR: Decoding JSON failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var request FindQuestionsBySurveyRequest
	if err := json.Unmarshal(in, &request); err != nil {
		log.Printf("ERROR: Unmarshal request failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	questions, err := h.service.FindQuestionsBySurvey(request.Name)
	if err != nil {
		log.Printf("ERROR: FindQuestionsBySurvey failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response FindQuestionsBySurveyResponse
	response.Questions = questions
	out, err := json.Marshal(response)
	if err != nil {
		log.Printf("ERROR: Marshal response failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

// InvokeFindQuestionsBySurvey ...
func InvokeFindQuestionsBySurvey(url string, request *FindQuestionsBySurveyRequest) (*FindQuestionsBySurveyResponse, error) {

	plain, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(base64.StdEncoding.EncodeToString(plain)))
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response FindQuestionsBySurveyResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
