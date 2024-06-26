package util

import (
	"fmt"
	"testing"
)

func ExampleReduce() {
	multiply := func(x, y int) int { return x * y }
	factors := List[int]([]int{1, 2, 3, 4})

	product := factors.Reduce(multiply)
	fmt.Println(product) // Output: 24
}

func ExampleFold() {
	multiplicand := 7
	multipliers := List[int]([]int{1, 2, 3, 4})
	multiply := func(x, y int) int { return x * y }

	product := multipliers.Fold(multiplicand, multiply)
	fmt.Println(product) // Output: 168
}

func TestList_Fold(t *testing.T) {
	t.Run("with a function reference", func(t *testing.T) {

	})

	t.Run("multiply integers", func(t *testing.T) {

	})

	t.Run("concatenate strings", func(t *testing.T) {

	})

	t.Run("with a closure capturing external scope", func(t *testing.T) {

	})

	t.Run("uncasted slice", func(t *testing.T) {

	})

	t.Run("empty list", func(t *testing.T) {

	})
}

func TestList_Reduce(t *testing.T) {
	t.Run("with an array", func(t *testing.T) {

	})

	t.Run("with pre-allocated slice", func(t *testing.T) {

	})

	t.Run("with a function reference", func(t *testing.T) {

	})

	t.Run("multiply integers", func(t *testing.T) {

	})

	t.Run("concatenate strings", func(t *testing.T) {

	})

	t.Run("uncasted slice", func(t *testing.T) {

	})

	t.Run("empty list - integer default value", func(t *testing.T) {

	})

	t.Run("empty list - string default value", func(t *testing.T) {

	})

	t.Run("empty list - struct default value", func(t *testing.T) {

	})
}
