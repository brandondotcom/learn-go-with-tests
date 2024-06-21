package sum

import (
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
