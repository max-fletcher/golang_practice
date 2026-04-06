package main

import "fmt"

// The 2 functions below were used before generics existed. The fact that you had to use 2 separate functions for the same operations because of
// difference in types was a nuisance.
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// After generics, this function can be used instead of 2 separate functions.
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	firstNums, secondNums := splitAnySlice(nums)
	fmt.Println(firstNums, secondNums)

	strs := []string{"henlo", "what", "are", "you", "doing"}
	firstStrs, secondStrs := splitAnySlice(strs)
	fmt.Println(firstStrs, secondStrs)
}
