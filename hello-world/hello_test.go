package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("it should sai Hello to name on param", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

		got2 := Hello("John")
		want2 := "Hello, John"

		if(got2 != want2) {
			t.Errorf("got: %q / want: %q", got2, want2)
		}
	})

	t.Run("It should say Hello, World if no params", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)		
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	
	if(got != want) {
		t.Errorf("got: %q / want: %q", got, want)
	}
}