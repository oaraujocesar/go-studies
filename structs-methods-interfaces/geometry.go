package geometry

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(r Rectangle) (perimeter float64) {
	return 2 * (r.width + r.height)
}

func Area(r Rectangle) (perimeter float64) {
	return r.width * r.height
}
