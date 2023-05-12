package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello to people when name supplied", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris!"

		if got != want {
			t.Errorf("should print %q but got %q", want, got)
		}
	})

	t.Run("Saying hello to all if name not supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"

		if got != want {
			t.Errorf("should print %q but got %q", want, got)
		}
	})

}
