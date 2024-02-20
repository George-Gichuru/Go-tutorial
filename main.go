package main

import "fmt"

func main() {
	// arrays

	var ages = [3]int{22, 23, 24}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var numbers = []int{2, 3, 4}
	numbers[2] = 25
	numbers = append(numbers, 10)
	fmt.Println(numbers, len(numbers))

	// slices ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
