package main

import (
	"fmt"
	"go-study/src/app"
	"io"
	"log"
	"net/http"
)

type MemoryPlayerStore struct {
	nameWithScore map[string]int
}

func (s *MemoryPlayerStore)GetScore(name string) int {
	return s.nameWithScore[name]
}

func main() {
	store := MemoryPlayerStore{
		map[string]int{"zsh": 20,
			"john": 10},
	}
	server := &app.PlayerServer{&store}

	err := http.ListenAndServe(":5000", server)
	if err != nil {
		log.Fatalf("listener error, %v", err)
	}
}

func countdown(writer io.Writer) {
	for i := 3;i > 0;i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
	fmt.Fprintf(writer, "GO")
}


