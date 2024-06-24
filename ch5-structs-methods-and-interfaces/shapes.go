package main

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the permiter of a rectangle given a height and width
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the area of a rectangle given the width and height
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
