package pwa

import (
	"log"
	"net/http"
)

// Serve ...
func Serve(address string) {
	for {
		log.Printf("INFO : Listening at [%s] ...", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalf("ERROR: Internal Server Error! [%s]", err.Error())
		}
	}
}
