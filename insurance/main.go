package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode int    `json:"zipcode"`
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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
