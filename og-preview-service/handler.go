package main

import (
	"io"
	"net/http"
)

func OGHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the body is closed after reading
	defer r.Body.Close()

	// Read the request body
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// If the body is empty, send a bad request error
	if len(data) == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}
	w.Write(data)
}
