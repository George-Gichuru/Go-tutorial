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
	fmt.Println(numbers)
}
