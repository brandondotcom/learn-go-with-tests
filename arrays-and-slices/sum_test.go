package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleReduceInt() {
	multiply := func(x, y int) int { return x * y }
	factors := []int{1, 2, 3, 4}

	product := ReduceInt(factors, multiply)
	fmt.Println(product) // Output: 24
}

func ExampleFoldInt() {
	multiplicand := 7
	multipliers := []int{1, 2, 3, 4}
	multiply := func(x, y int) int { return x * y }

	product := FoldInt(multipliers, multiplicand, multiply)
	fmt.Println(product) // Output: 168
}

func TestMult(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Mult(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}

		got := Mult(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		got := Mult(nil)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all nested slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		checkAllSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkAllSums(t, got, want)
	})

	t.Run("sum all tails with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{})
		want := []int{2, 9, 0}

		checkAllSums(t, got, want)
	})
}

func checkAllSums(t *testing.T, got []int, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
