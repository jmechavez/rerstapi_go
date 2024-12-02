package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", allClient).Methods(http.MethodGet)
	router.HandleFunc("/clients", createClient).Methods(http.MethodPost)

	router.HandleFunc("/clients/{client_id:[0-9]+}", getClients).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
