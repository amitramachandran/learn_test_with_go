package iterators

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output : "aaaaaa"
}

func TestRepeat(t *testing.T) {
	repeat_number := 5
	result := Repeat("go", repeat_number)
	expected := "gogogogogo"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
