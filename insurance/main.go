package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Client struct {
	Name    string `xml:"name"`
	City    string `xml:"city"`
	Zipcode int    `xml:"zipcode"`
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/clients", allClient)

	// starting server
	http.ListenAndServe("localhost:8000", nil)
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
