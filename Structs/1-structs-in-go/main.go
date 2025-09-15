package main

import (
	"fmt"
)

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	// Structs are like JS objects

	message := messageToSend{}

	message.phoneNumber = 897908905678
	message.message = "Henlo"

	fmt.Println(message)
}
