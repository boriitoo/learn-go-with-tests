package context

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response 	string
	cancelled 	bool
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	want := "hello world!"

	server := Server(&SpyStore{want})

	request := httptest.NewRequest(http.MethodGet, "/", nil);
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
