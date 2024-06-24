package main

import "math"

type Shape interface {
	Area() float64
	GetShape() Shape
}

type Squareable interface {
	MakeSquare()
}

func (r *Rectangle) MakeSquare() {
	r.Width = r.Height
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) GetShape() Shape {
	return r
}

func (r Circle) GetShape() Shape {
	return r
}

func (r Triangle) GetShape() Shape {
	return r
}

var _ Squareable = (*Rectangle)(nil)
var _ Squareable = &Rectangle{}
var _ Shape = Rectangle{}
var _ Shape = (*Rectangle)(nil)

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

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	return 0
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
