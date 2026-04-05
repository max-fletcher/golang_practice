package main

import (
	"fmt"
	"time"
)

// Basically a goroutine that can be terminated by the parent(main). default contains the logic that needs to be executed ultil the parent wants to cancel it.
// If we close this channnel, the 1st case statement will be triggered instead of the default and the return statement will cause the goroutine
// to terminate.
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("...Doing work")
		}
	}
}

func main() {
	// 1st example(not practical at all)
	// fmt.Println("Hi from main")

	// channel1 := make(chan string, 3)
	// chars := []string{"a", "b", "c"}

	// // send chars into the channel
	// for _, s := range chars {
	// 	select {
	// 	case channel1 <- s:
	// 	}
	// }

	// close(channel1)

	// // receive data from the channel and print it
	// for result := range channel1 {
	// 	fmt.Println(result)
	// }

	// 2nd example
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)

	return
}
