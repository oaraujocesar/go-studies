package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{width: 10.0, height: 10.0}

	received := Perimeter(r)
	expected := 40.0

	if received != expected {
		t.Errorf("received %.2f but expected %.2f", received, expected)
	}
}

func TestArea(t *testing.T) {
	r := Rectangle{width: 5.0, height: 15.0}
	received := Area(r)
	expected := 75.0

	if received != expected {
		t.Errorf("received %.2f but expected %.2f", received, expected)
	}
}
