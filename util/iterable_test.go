package util

import "fmt"

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
