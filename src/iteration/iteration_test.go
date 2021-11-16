package iteration

import "testing"

func TestRepeat(t *testing.T) {
	str := repeat("a", 5)
	expected := "aaaaa"
	if str != expected {
		t.Errorf("expected '%s' but got '%s'", expected, str)
	}
}
