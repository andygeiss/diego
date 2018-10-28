package main

import (
	"github.com/andygeiss/diego/pkg/pwa"
	"log"
	"os"
)

func main() {

	bindAddr := os.Getenv("BIND")
	engineURL := os.Getenv("ENGINE_URL")
	surveyName := os.Getenv("SURVEY")
	title := os.Getenv("APP_TITLE")

	log.Printf("INFO : Bind address is  [%s]", bindAddr)
	log.Printf("INFO : Engine URL is    [%s]", engineURL)
	log.Printf("INFO : App title is     [%s]", title)
	log.Printf("INFO : Survey name is   [%s]", surveyName)
	log.Printf("INFO : Serving App ...")

	pwa.Init()
	pwa.AddPage("/", "web/templates/index.html", engineURL, surveyName, title)
	pwa.Serve(bindAddr)
}
