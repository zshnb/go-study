package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	countdown(&buffer)
	actual := buffer.String()
	expected := `3
2
1
GO`
	if actual != expected {
		t.Errorf("expect %s but actual %s", expected, actual)
	}
}
