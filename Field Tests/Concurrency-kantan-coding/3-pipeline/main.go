package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	outChannel := make(chan int)
	go func() {
		for _, n := range nums {
			outChannel <- n
		}
		close(outChannel)
	}()

	return outChannel
}

func sq(inChannel <-chan int) <-chan int {
	outChannel := make(chan int)

	go func() {
		for n := range inChannel {
			outChannel <- n * n
		}
		close(outChannel)
	}()

	return outChannel
}

func main() {
	// input
	nums := []int{1, 2, 3, 4, 5, 6}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage2
	finalChannel := sq(dataChannel)

	for num := range finalChannel {
		fmt.Println(num)
	}

	return
}
