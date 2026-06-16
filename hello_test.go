package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"

		assertCorrectMessage(got, want, t)
	})
	t.Run("say \"Hello, world\" when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(got, want, t)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("James", "Spanish")
		want := "Hola, James"

		assertCorrectMessage(got, want, t)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Arno", "French")
		want := "Bonjour, Arno"

		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t *testing.T) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
