package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}

	received := Perimeter(rectangle)
	expected := 40.0

	if received != expected {
		t.Errorf("received %.2f but expected %.2f", received, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2.0, Height: 15.0}, hasArea: 34.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			received := tt.shape.Area()

			if received != tt.hasArea {
				t.Errorf("%#v Received %g but expected %g", tt.shape, received, tt.hasArea)
			}
		})
	}
}
