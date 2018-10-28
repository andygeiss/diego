package handlers

import (
	"github.com/andygeiss/diego/pkg/render"
	"log"
	"net/http"
)

// StaticPageHandler ...
type StaticPageHandler struct {
	content []byte
}

// NewStaticPageHandler ...
func NewStaticPageHandler(templateFile, engineURL, surveyName, title, scripts, styles string) http.Handler {
	return &StaticPageHandler{
		content: embed(templateFile, engineURL, surveyName, title, scripts, styles),
	}
}

// ServeHTTP serves the current content to the client.
func (h *StaticPageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write(h.content)
}

// embed creates the static content by using the template engine.
func embed(templateFile, engineURL, surveyName, title, scripts, styles string) []byte {
	out, err := render.Template(templateFile, struct {
		EngineURL  string
		Scripts    string
		Styles     string
		SurveyName string
		Title      string
	}{
		EngineURL:  engineURL,
		Scripts:    scripts,
		Styles:     styles,
		SurveyName: surveyName,
		Title:      title,
	})
	if err != nil {
		log.Printf("ERROR: Rendering of [%s] failed! [%s]", templateFile, err.Error())
	}
	return out
}
