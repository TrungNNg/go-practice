package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOGHandler(t *testing.T) {

	t.Run("parse url from json", func(t *testing.T) {
		inputJSON := `{"url":"https://www.imdb.com/title/tt27987046/"}`
		request := newPostRequest(inputJSON)
		response := httptest.NewRecorder()

		OGHandler(response, request)

		// result from response recorder
		rs := response.Result()

		got, _ := io.ReadAll(rs.Body)
		defer rs.Body.Close()
		want := inputJSON

		assertEqual(t, string(got), want)

		// check for header Content-Type: application/json
		wantHeaderValue := "application/json"
		assertEqual(t, rs.Header.Get("Content-Type"), wantHeaderValue)

		// check status code
		assertEqual(t, rs.Status, "200 OK")
	})

}

// this handler will be passed into a httptest.Server, so that for different
// path the server will return different html
func mockHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/h1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head><title>Mock Page</title></head><body><h1>Hello, World!</h1></body></html>`))
	})
	mux.HandleFunc("/h2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html><head>Heelo</head></html>`))
	})
	return mux
}

func TestFetchHTML(t *testing.T) {
	ms := mockServer()
	got, _ := fetchHTML(ms.URL + "/h1")
	want := `<html><head><title>Mock Page</title></head><body><h1>Hello, World!</h1></body></html>`

	assertEqual(t, got, want)

	got, _ = fetchHTML(ms.URL + "/h2")
	want = `<html><head>Heelo</head></html>`

	assertEqual(t, got, want)
}

func mockServer() *httptest.Server {
	return httptest.NewServer(mockHandler())
}

func newPostRequest(json string) *http.Request {
	// Convert the string to an io.Reader
	body := strings.NewReader(json)
	request, _ := http.NewRequest(http.MethodPost, "/og", body)
	return request
}

// should be change to assertResponseJSON later
func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, \ngot  %q \nwant %q", got, want)
	}
}
