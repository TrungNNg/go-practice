package main

import (
	"net/http"
)

func main() {
	handler := http.HandlerFunc(OGHandler)
	http.ListenAndServe(":5000", handler)
}
