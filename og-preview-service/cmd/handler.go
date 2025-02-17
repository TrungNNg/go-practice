package main

import (
	"encoding/json"
	"net/http"
)

func OGHandler(w http.ResponseWriter, r *http.Request) {
	// struct to hold json input
	var input struct {
		URL string `json:"url"`
	}

	// Refactor : this should be a separate func
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// use the url and make the get request to get the html, parse the html to return go
	// object

	// Refactor: separate func
	data := map[string]string{
		"url": input.URL,
	}

	js, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// separate func ?
	w.Header().Set("Content-Type", "application/json")

	// separate func to write json only
	w.Write(js)

}
