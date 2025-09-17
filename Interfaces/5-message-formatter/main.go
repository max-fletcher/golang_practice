package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}

func (p plainText) format() string {
	return p.message
}

type bold struct {
	message string
}

func (b bold) format() string {
	return fmt.Sprintf("**%s**", b.message)
}

type code struct {
	message string
}

func (c code) format() string {
	return fmt.Sprintf("`%s`", c.message)
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}

func main() {
	plText := plainText{
		message: "Plaintext boi",
	}

	fmt.Println(sendMessage(plText))

	bText := bold{
		message: "Bold boi",
	}

	fmt.Println(sendMessage(bText))

	cText := code{
		message: "Code boi",
	}

	fmt.Println(sendMessage(cText))
}
