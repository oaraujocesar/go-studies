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

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{width: 2.0, height: 15.0}, 34.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		received := tt.shape.Area()

		if received != tt.expected {
			t.Errorf("Received %g but expected %g", received, tt.expected)
		}
	}
}
