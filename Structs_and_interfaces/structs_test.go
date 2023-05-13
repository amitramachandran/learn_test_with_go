package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{4.0, 5.0}
	got := Perimeter(r)
	expected := 18.0

	if got != expected {
		t.Errorf("got %.2f but expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{4.0, 5.0}, 20.0},
		{"circle", Circle{4.0}, 50.27},
		{"triangle", Triangle{4.0, 5.0}, 10.0},
	}

	for _, sh := range areaTests {
		t.Run(sh.name, func(t *testing.T) {
			got := sh.shape.Area()
			if got != sh.expected {
				t.Errorf("%#v got %.2f but expected %.2f", sh.name, got, sh.expected)
			}
		})

	}
}

func TestAreaInterface(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got %.2f but expected %.2f", got, expected)
		}
	}
	t.Run("area test for rectangle", func(t *testing.T) {
		r := Rectangle{4.0, 5.0}
		expected := 20.0
		checkArea(t, r, expected)

	})

	t.Run("Area test for circle", func(t *testing.T) {
		c := Circle{4.0}
		expected := 50.27
		checkArea(t, c, expected)

	})

}
