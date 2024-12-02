package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Client struct {
	Name    string `xml:"name"    json:"name"`
	City    string `xml:"city"    json:"city"`
	Zipcode int    `xml:"zipcode" json:"zipcode"`
}

func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post a request")
}

func getClients(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["client_id"])
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func allClient(w http.ResponseWriter, r *http.Request) {
	clients := []Client{
		{
			Name:    "John Michael",
			City:    "Gensan",
			Zipcode: 9500,
		},
		{
			Name:    "Jm",
			City:    "Paranaque",
			Zipcode: 1700,
		},
	}

	// if xml format if not json format
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(clients)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)
	}
}
