package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	mail := email{
		message:     "New msg 1",
		fromAddress: "From address 1",
		toAddress:   "To address 1",
	}

	mail.setMessage("New msg boi 1")
	fmt.Println("Mail:", mail)
}
