package dependency_inject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet(&buffer, "zsh")
	actual := buffer.String()
	expected := "hello zsh"
	if actual != expected {
		t.Errorf("expect %s but actual %s", expected, actual)
	}
}