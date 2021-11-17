package maps

import "testing"

func assertString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expect %s but actual %s", expected, actual)
	}
}
func TestSearch(t *testing.T) {
	t.Run("contains key", func(t *testing.T) {
		dict := Dictionary{"name": "zsh"}
		expected := "zsh"
		actual, _ := dict.search("name")
		assertString(t, expected, actual)
	})

	t.Run("not contains key", func(t *testing.T) {
		dict := Dictionary{"name": "zsh"}
		expected := "key not exist"
		_, error := dict.search("key")
		if error == nil {
			t.Fatal("expected error")
		}
		assertString(t, expected, error.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add not exist key", func(t *testing.T) {
		dict := Dictionary{}
		dict.add("name", "zsh")
		actual, _ := dict.search("name")
		expected := "zsh"
		assertString(t, expected, actual)
	})
}