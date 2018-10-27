package main

import (
	"github.com/andygeiss/diego/pkg/pwa"
	"log"
)

func main() {
	log.Printf("INFO : Serving App ...")
	pwa.Init()
	pwa.AddPage("/", "web/templates/index.html", "App")
	pwa.Serve(":80")
}
