package dependencyinjection

import (
	"bytes"
	"testing"
)

// we are using bytes and buffer interface to send the data
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")
	got := buffer.String()
	want := "Hello Chris"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
