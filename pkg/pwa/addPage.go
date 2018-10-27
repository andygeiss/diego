package pwa

import (
	"encoding/base64"
	"github.com/andygeiss/diego/pkg/handlers"
	"github.com/andygeiss/diego/pkg/parse"
	"net/http"
)

// AddPage ...
func AddPage(pattern, templateFile, title string) {
	styles := base64.StdEncoding.EncodeToString(parse.Files("web/css/styles.css"))
	scripts := base64.StdEncoding.EncodeToString(parse.Files("vendor/wasm_exec.js", "web/js/init.js"))
	http.Handle(pattern, handlers.NewStaticPageHandler(templateFile, title, scripts, styles))
}
