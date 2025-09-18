package main

import "fmt"

func countConnections(groupSize int) int {
	return (groupSize * (groupSize - 1)) / 2
}

func main() {
	count := countConnections(5)
	fmt.Println(count)
}
