package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/andygeiss/diego/internal/survey"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type getResultsByFactsHandler struct {
	service survey.Service
}

// GetResultsByFactsRequest ...
type GetResultsByFactsRequest struct {
	Facts []string `json:"facts"`
}

// GetResultsByFactsResponse ...
type GetResultsByFactsResponse struct {
	Results []string `json:"results"`
}

// NewGetResultsByFactsHandler ...
func NewGetResultsByFactsHandler(service survey.Service) http.Handler {
	return &getResultsByFactsHandler{service: service}
}

// ServeHTTP ...
func (h *getResultsByFactsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	var request GetResultsByFactsRequest
	if err := json.Unmarshal(in, &request); err != nil {
		log.Printf("ERROR: Unmarshal request failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	results, err := h.service.GetResultsByFacts(request.Facts)
	if err != nil {
		log.Printf("ERROR: GetResultsByFacts failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response GetResultsByFactsResponse
	response.Results = results
	out, err := json.Marshal(response)
	if err != nil {
		log.Printf("ERROR: Marshal response failed! [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

// InvokeGetResultsByFacts ...
func InvokeGetResultsByFacts(url string, request *GetResultsByFactsRequest) (*GetResultsByFactsResponse, error) {

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

	var response GetResultsByFactsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
