package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockCheck(url string) bool {
	if url == "a" {
		return true
	}
	return false
}

func slowCheck(url string) bool  {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsiteCheck(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100;i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		checkWebsites(slowCheck, urls)
	}
}

func TestWebsiteCheck(t *testing.T) {
	expected := map[string]bool{
		"a": true,
		"b": false,
	}
	urls := []string{"a", "b"}
	actual := checkWebsites(mockCheck, urls)
	if len(expected) != len(actual) {
		t.Errorf("length inconsistent")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("map inconsistent")
	}
}
