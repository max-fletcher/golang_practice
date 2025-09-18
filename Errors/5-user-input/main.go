package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}

	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}

	return errors.New(status)
}

func main() {
	err := validateStatus("Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo Henlo ")

	fmt.Println(err)
}
