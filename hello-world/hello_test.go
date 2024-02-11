package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("it should sai Hello to name on param", func(t *testing.T) {
		got := Hello("Chris", "EN")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

		got2 := Hello("John", "EN")
		want2 := "Hello, John"

		if(got2 != want2) {
			t.Errorf("got: %q / want: %q", got2, want2)
		}
	})

	t.Run("It should say Hello, World if no params", func(t *testing.T) {
		got := Hello("", "EN")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)		
	})

	t.Run("It should greet in Spanish, and then in English", func(t *testing.T) {
		got := Hello("Joao", "ES")
		want := "Hola, Joao"
		assertCorrectMessage(t, got, want)

		got = Hello("John", "EN")
		want = "Hello, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("If not mapped language, should default to English", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello, Bob"

		assertCorrectMessage(t, got, want)

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	
	if(got != want) {
		t.Errorf("got: %q / want: %q", got, want)
	}
}