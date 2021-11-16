package hello

import "testing"

func assertEquals(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("expect %q but actual %q", want, got)
	}
}

func TestHello(t *testing.T) {
	t.Run("say something", func(t *testing.T) {
		got := getString("zsh", "english")
		want := "hello zsh"
		assertEquals(t, got, want)
	})
	t.Run("say empty", func(t *testing.T) {
		got := getString("", "english")
		want := "hello world"
		assertEquals(t, got, want)
	})
	t.Run("say in spanish", func(t *testing.T) {
		got := getString("zsh", "spanish")
		want := "hola zsh"
		assertEquals(t, got, want)
	})
	t.Run("say in french", func(t *testing.T) {
		got := getString("zsh", "french")
		want := "bonjue zsh"
		assertEquals(t, got, want)
	})
}
