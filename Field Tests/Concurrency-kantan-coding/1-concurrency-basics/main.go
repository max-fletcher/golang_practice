package main

import (
	"fmt"
	"time"
)

// Start: Defining enums
type OrderStatus string

// const (
// 	StatusPending   OrderStatus = "pending"
// 	StatusShipped   OrderStatus = "shipped"
// 	StatusDelivered OrderStatus = "delivered"
// 	StatusCancelled OrderStatus = "cancelled"
// )

func someFunc(str string) {
	fmt.Printf("Hello from somefunc: %v\n", str)
}

func main() {
	fmt.Println("Hi from main")

	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(time.Second * 1)

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "data boi 1"
	}()

	go func() {
		channel2 <- "data boi 2"
	}()

	select {
	case msgFromChannel1 := <-channel1:
		fmt.Println(msgFromChannel1)
	case msgFromChannel2 := <-channel2:
		fmt.Println(msgFromChannel2)
	}

	return
}
