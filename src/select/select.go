package _select

import (
	"net/http"
	"time"
)

func race(url1, url2 string) string {
	startA := time.Now()
	http.Get(url1)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(url2)
	durationB := time.Since(startB)

	if durationA < durationB {
		return url1
	}
	return url2
}
