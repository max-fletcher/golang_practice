package main

import "fmt"

// nested structs i.e one struct can contain another struct(i.e like in TS nested type)
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 || mToSend.recipient.name == "" || mToSend.recipient.number == 0 || mToSend.message == "" {
		return false
	}

	return true
}

func main() {
	// Structs are like JS objects
	// Create a new struct with default values
	message := messageToSend{}
	message.sender.name = "Max"
	message.sender.number = 2399021839021

	message.recipient.name = "Fletcher"
	message.recipient.number = 897908905678

	message.message = "Henlo"

	fmt.Println(message)
}
