package sum

import "learn-go-with-tests/ch2-integers"

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func Tail(numbers []int) (tail []int) {
	if len(numbers) == 0 {
		tail = numbers
	} else {
		tail = numbers[1:]
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(Tail(numbers)))
	}

	return
}

// Multiplies numbers in array, 0 if empty
func Mult(numbers []int) int {
	multiply := func(x, y int) int {
		return x * y
	}

	return ReduceInt(numbers, multiply)
}

func Sum(numbers []int) int {
	return ReduceInt(numbers, integers.Add)
}

type IntOperation func(acc, next int) int

// FoldInt Accumulates value starting w/ initial and applying operation from left to right
func FoldInt(numbers []int, initial int, operation IntOperation) int {
	if len(numbers) == 0 {
		return initial
	} else {
		return FoldInt(numbers[1:], operation(initial, numbers[0]), operation)
	}
}

// ReduceInt Accumulates value applying operation from left to right
func ReduceInt(numbers []int, operation IntOperation) int {
	if len(numbers) == 0 {
		return 0
	} else {
		return FoldInt(numbers[1:], numbers[0], operation)
	}
}
