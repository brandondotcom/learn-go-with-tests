package util

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
