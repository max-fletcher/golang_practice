package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages)) // make new slice

	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}

	return costs
}

func main() {
	arr := []string{"I", "hate", "my", "life", "^_^"}
	prices := getMessageCosts(arr)
	fmt.Println(prices)
}
