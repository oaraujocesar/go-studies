package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) (perimeter float64) {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	pi := math.Pi
	r2 := math.Pow(c.Radius, 2)

	return r2 * pi
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
