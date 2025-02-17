package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

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
	req, err := http.NewRequest("GET", input.URL, nil)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	//data2, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(data2))
	defer resp.Body.Close()

	// parse html

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
