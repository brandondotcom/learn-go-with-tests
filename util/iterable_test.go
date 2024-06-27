package util

import (
	"fmt"
	"slices"
	"testing"
)

func ExampleList_Reduce() {
	multiply := func(x, y int) int { return x * y }
	factors := List[int]([]int{1, 2, 3, 4})

	product := factors.Reduce(multiply)
	fmt.Println(product) // Output: 24
}

func ExampleList_Fold() {
	multiplicand := 7
	multipliers := List[int]([]int{1, 2, 3, 4})
	multiply := func(x, y int) int { return x * y }

	product := multipliers.Fold(multiplicand, multiply)
	fmt.Println(product) // Output: 168
}

func TestNewList(t *testing.T) {
	assertEqual := func(got, want []int, t *testing.T) {
		t.Helper()
		if !slices.EqualFunc(got, want, func(x, y int) bool { return x == y }) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}

	t.Run("with trailing arguments", func(t *testing.T) {
		got := NewList(0, 1, 2, 3)
		want := []int{0, 1, 2, 3}

		assertEqual(got, want, t)
	})

	t.Run("with array expansion", func(t *testing.T) {
		want := []int{0, 1, 2, 3}
		got := NewList(want...)

		assertEqual(got, want, t)
	})
}

func TestList_Fold(t *testing.T) {

	t.Run("with a function reference", func(t *testing.T) {

	})

	t.Run("with a closure capturing local scope", func(t *testing.T) {

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
