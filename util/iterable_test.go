package util

import "fmt"

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
