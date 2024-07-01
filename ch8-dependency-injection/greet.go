package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// greeter holds the config/dependencies for the Greet function.
type greeter struct {
	writer io.Writer
	name   string
}

func (g *greeter) greet() {
	fmt.Fprintf(g.writer, "Hello, %s", g.name)
}

// greetOption is any function that takes a single argument reference to a greeter.
// It can be used as a return value for functions that implement the Functional Options pattern.
// see: https://golang.cafe/blog/golang-functional-options-pattern.html
type greetOption func(*greeter)

func (g *greeter) setDefaults() {
	g.writer = os.Stdout
	g.name = "World"
}

// region options

func WithWriter(writer io.Writer) greetOption {
	return func(g *greeter) {
		g.writer = writer
	}
}

func WithName(name string) greetOption {
	return func(g *greeter) {
		g.name = name
	}
}

// endregion

func Greet(options ...greetOption) {
	g := &greeter{}
	g.setDefaults()

	for _, o := range options {
		o(g)
	}

	g.greet()
}

// region main

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(
		WithWriter(w),
		WithName("world"),
	)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

// endregion
