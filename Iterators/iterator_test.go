package iterators

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("go")
	expected := "gogogogogo"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
