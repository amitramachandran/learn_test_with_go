package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Hello"
	misc := func(string) {}
	type x struct {
		firstname string
		lastname  string
	}
	X := x{"Hello", "kumar"}
	got := Walk(X, misc)
	want := expected

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}
