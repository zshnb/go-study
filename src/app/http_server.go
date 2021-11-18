package app

import (
	"fmt"
	"net/http"
)
type PlayerStore interface {
	GetScore(name string) int
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		player := r.URL.Path[len("/player/"):]
		score := p.Store.GetScore(player)
		if score == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		fmt.Fprintf(w, "%d", score)
	} else if r.Method == http.MethodPost {
		
	}
}
