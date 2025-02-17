package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestOGHander(t *testing.T) {

	t.Run("return string", func(t *testing.T) {
		inputJSON := `
		{
			"url":"https://www.imdb.com/title/tt27987046/"
		}
		`
		request := newPostRequest(inputJSON)
		response := httptest.NewRecorder()

		OGHandler(response, request)

		got := response.Body.String()
		want := inputJSON

		assertResponseBody(t, got, want)
	})

}

func newPostRequest(json string) *http.Request {
	// Convert the string to an io.Reader
	body := strings.NewReader(json)
	request, _ := http.NewRequest(http.MethodPost, "/og", body)
	return request
}

// should be change to assertResponseJSON later
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
