package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planFree {
		return messages[0:2], nil
	}
	if plan == planPro {
		return messages[:], nil
	}

	return nil, errors.New("unsupported plan")
}

func main() {
	messages := [3]string{"Henlo", "Worl", "Boi"}
	msgs, err := getMessageWithRetriesForPlan(planFree, messages)
	fmt.Println(msgs, err)

	// messages1 := [6]string{"Henlo", "Worl", "Boi", "Lololol", "Wee", "Woo"}
	// messages1[2] = "Hoo"
	// fmt.Println(messages1[1:3])
}
