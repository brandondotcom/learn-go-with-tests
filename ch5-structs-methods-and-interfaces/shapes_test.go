package main

import (
	"reflect"
	"testing"
)

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
		{name: "Rectangle Value", shape: Rectangle{Width: 12, Height: 6}, isSquareable: false /*why?*/},
		{name: "Rectangle Reference", shape: &Rectangle{Width: 12, Height: 6}, isSquareable: true},
		{name: "Circle", shape: &Circle{Radius: 10}, isSquareable: false},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, isSquareable: false},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.shape.(type) {
			default:
				t.Logf("type of %s, %T", tt.name, v)
				t.Logf("type of %s, %v", tt.name, reflect.TypeOf(tt.shape))
			}
			_, squareableCheck := tt.shape.(Squareable)
			if squareableCheck != tt.isSquareable {
				t.Errorf("got %t when checking %s for Squareable, expected %t", squareableCheck, tt.name, tt.isSquareable)
			}

			_, shapeCheck := tt.shape.(Shape)

			if !shapeCheck {
				t.Errorf("expected %s to implement Shape, but didn't", tt.name)
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
