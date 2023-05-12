package main

import "testing"

func AssertCorrectGreet(got string, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("should print %q but got %q", want, got)
	}
}

func TestHello(t *testing.T) {

	t.Run("Saying hello to people when name supplied", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris!"
		AssertCorrectGreet(got, want, t)

	})

	t.Run("Saying hello to all if name not supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		AssertCorrectGreet(got, want, t)
	})

	t.Run("Greeting with different language", func(t *testing.T) {
		got := Hello("", "french")
		want := "Hola World!"
		AssertCorrectGreet(got, want, t)
	})

}
