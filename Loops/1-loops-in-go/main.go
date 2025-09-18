package main

import "fmt"

// func bulkSend(numMessages int) float64 {
// 	inc := 0.0
// 	total := float64(numMessages)
// 	for i := 1; i <= numMessages; i++ {
// 		total = total + inc
// 		inc = inc + 0.01
// 	}

// 	return total
// }

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}

	return totalCost
}

func main() {
	totalCost := bulkSend(10)

	fmt.Println(totalCost)
}
