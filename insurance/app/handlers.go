package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/jmechavez/restapi_go/insurance/service"
)

type ClientHandlers struct {
	service service.ClientService
}

func (ch *ClientHandlers) getAllClient(w http.ResponseWriter, r *http.Request) {
	clients, _ := ch.service.GetAllClient()
	// if xml format if not json format
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(clients)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}
