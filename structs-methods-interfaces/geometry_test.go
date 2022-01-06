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

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{width: 2.0, height: 15.0}

		received := rectangle.Area()
		expected := 34.0

		if received != expected {
			t.Errorf("received %g but expected %g", received, expected)
		}
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}

		received := circle.Area()
		expected := 314.1592653589793

		if received != expected {
			t.Errorf("received %g but expected %g", received, expected)
		}
	})
}
