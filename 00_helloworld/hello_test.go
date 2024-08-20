package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello Bob"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' if empty string is supplied'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Alberto", "Spanish")
		want := "Hola Alberto"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say in French", func(t *testing.T) {
		got := Hello("Bob", "French")
		want := "Bonjour Bob"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\nGot:  %q\nWant: %q", got, want)
	}
}
