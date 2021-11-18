package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	actual := race(slowUrl, fastUrl)
	if actual != fastUrl {
		t.Errorf("expect %s but actual %s", fastUrl, actual)
	}

	defer slowServer.Close()
	defer fastServer.Close()
}