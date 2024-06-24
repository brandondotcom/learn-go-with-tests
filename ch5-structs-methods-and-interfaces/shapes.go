package main

// Perimeter calculates the permiter of a rectangle given a height and width
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Area calculates the area of a rectangle given the width and height
func Area(width float64, height float64) float64 {
	return width * height
}
