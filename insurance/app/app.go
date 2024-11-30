package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/clients", allClient)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
