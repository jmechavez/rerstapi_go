package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/restapi_go/insurance/domain"
	"github.com/jmechavez/restapi_go/insurance/service"
)

func Start() {
	router := mux.NewRouter()

	ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryStub())}

	// define routes
	router.HandleFunc("/clients", ch.getAllClient).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
