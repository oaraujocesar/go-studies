package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 10.0}

	received := Perimeter(rectangle)
	expected := 40.0

	if received != expected {
		t.Errorf("received %.2f but expected %.2f", received, expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()

		received := shape.Area()

		if received != expected {
			t.Errorf("Received %g but Expected %g", received, expected)
		}
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{width: 2.0, height: 15.0}

		checkArea(t, rectangle, 34.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
}
