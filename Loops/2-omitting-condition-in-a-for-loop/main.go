package main

import "fmt"

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (float64(i) * 0.01)
		if totalCost > thresh {
			return i
		}
	}
}

func main() {
	maxMsg := maxMessages(20)

	fmt.Println(maxMsg)
}
