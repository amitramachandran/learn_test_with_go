package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output : 9
}

func TestAdd(t *testing.T) {
	var x, y int
	x = 2
	y = 3
	got := Add(x, y)
	want := 5

	if got != want {
		t.Errorf("The add function must return %d but got %d", want, got)
	}
}
