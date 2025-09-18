package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	msg1, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	msg2, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return msg1 + msg2, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	totalCost, err := sendSMSToCouple("Henlo Henlo Henlo Henlo", "Boi Boi Boi Boi")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total Cost: ", totalCost)
}
