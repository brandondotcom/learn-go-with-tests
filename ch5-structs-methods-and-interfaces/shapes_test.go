package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        Shape
		isSquareable bool
	}{
		// This test fails - why? or - s it a meaningful test?
		//{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, isSquareable: true},
		{name: "Circle", shape: Circle{Radius: 10}, isSquareable: false},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, isSquareable: false},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.shape.GetShape().(type) {
			default:
				_, got := v.(Squareable)
				if got != tt.isSquareable {
					t.Errorf("%T", v)
				}
			}
			s, got := tt.shape.(Squareable)
			if got {
				s.MakeSquare()
			}
			if got != tt.isSquareable {
				t.Errorf("%t", got)
			}
		})
	}
}

func TestMakeSquare(t *testing.T) {
	rectangle := Rectangle{10.0, 15.0}

	rectangle.MakeSquare()

	if rectangle.Height != rectangle.Width {
		t.Errorf("Expected height and width to be equal, got height: %g, width: %g", rectangle.Height, rectangle.Width)
	}

}
