package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Matan", "")
		want := "Hello, Matan"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Roxane", "French")
		want := "Bonjour, Roxane"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Hebrew", func(t *testing.T) {
		got := Hello("Matan", "Hebrew")
		want := "Shalom, Matan"

		assertCorrectMessage(t, got, want)
	})
}
