package pwa

import (
	"github.com/andygeiss/diego/pkg/handlers"
	"github.com/andygeiss/diego/pkg/parse"
	"net/http"
)

func Init() {
	http.Handle("/lib.wasm", handlers.NewStaticDataHandler(parse.Files("build/package/lib.wasm"), "application/wasm"))
	http.Handle("/favicon.ico", handlers.NewStaticDataHandler(parse.Files("web/img/favicon.ico"), "image/x-icon"))
	http.Handle("/img/logo.png", handlers.NewStaticDataHandler(parse.Files("web/img/logo.png"), "image/png"))
	http.Handle("/manifest.json", handlers.NewStaticDataHandler(parse.Files("web/manifest.json"), "application/json"))
	http.Handle("/service-worker.js", handlers.NewStaticDataHandler(parse.Files("web/js/service-worker.js"), "application/javascript"))
}
