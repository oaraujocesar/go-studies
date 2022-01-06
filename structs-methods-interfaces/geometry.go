package geometry

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) (perimeter float64) {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	pi := math.Pi
	r2 := math.Pow(c.Radius, 2)

	return r2 * pi
}
