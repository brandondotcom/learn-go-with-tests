package sum

import (
	"learn-go-with-tests/ch2-integers"
	"learn-go-with-tests/util"
)

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

	return util.ReduceInt(numbers, multiply)
}

func Sum(numbers []int) int {
	return util.ReduceInt(numbers, integers.Add)
}
