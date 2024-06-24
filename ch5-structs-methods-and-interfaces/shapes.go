package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimiter of a rectangle given a height and width
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a rectangle given the width and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * (c.Radius)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
