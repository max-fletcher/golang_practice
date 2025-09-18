package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

func main() {
	// Adding an index inside [] will throw an error because variadic functions by definition cannot have a fixed length, and so we can't pass an array; only slices will be accepted.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	count := sum(nums...)
	fmt.Println(count)
}
