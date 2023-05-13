package structs

import "math"

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	s_width  float64
	s_height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	s_base   float64
	s_height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.s_width + r.s_height)
}

func (r Rectangle) Area() float64 {
	return r.s_height * r.s_width
}

func (c Circle) Area() float64 {
	val := math.Pi * math.Pow(c.radius, 2)
	area := roundFloat(val, 2)
	return area
}

func (t Triangle) Area() float64 {
	val := 0.5 * t.s_base * t.s_height
	return val
}
