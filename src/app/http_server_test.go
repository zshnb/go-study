package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	nameWithScore map[string]int
}

func (s *StubPlayerStore)GetScore(name string) int {
	return s.nameWithScore[name]
}

func TestServer(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{"zsh": 20,
			"john": 10},
	}
	t.Run("return zsh's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/zsh", nil)
		response := httptest.NewRecorder()
		server := &PlayerServer{&store}
		server.ServeHTTP(response, request)

		actual := response.Body.String()
		expected := "20"
		if actual != expected {
			t.Errorf("expect %s but actual %s", expected, actual)
		}
	})
	t.Run("return john's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/john", nil)
		response := httptest.NewRecorder()
		server := &PlayerServer{&store}
		server.ServeHTTP(response, request)

		actual := response.Body.String()
		expected := "10"
		if actual != expected {
			t.Errorf("expect %s but actual %s", expected, actual)
		}
	})
	t.Run("return 404 when user not found", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/player/abc", nil)
		response := httptest.NewRecorder()
		server := &PlayerServer{&store}
		server.ServeHTTP(response, request)

		actual := response.Code
		expected := http.StatusNotFound
		if actual != expected {
			t.Errorf("expect %d but actual %d", expected, actual)
		}
	})
}
