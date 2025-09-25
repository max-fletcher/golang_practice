package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}

	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

type testCase struct {
	messageIn *string
	expected  *string
}

func main() {
	s1 := "English, motherfubber, do you speak it?"
	s2 := "English, mother****er, do you speak it?"
	s3 := "Does he look like a witch?"
	s4 := "Does he look like a *****?"

	testCases := []testCase{
		{
			&s1,
			&s2,
		},
		{
			nil,
			nil,
		},
		{
			&s3,
			&s4,
		},
		{
			nil,
			nil,
		},
	}

	for i := range testCases {
		removeProfanity(testCases[i].messageIn)

		if testCases[i].messageIn != nil && testCases[i].expected != nil { // For when pointers are pointing to a number
			if *testCases[i].messageIn == *testCases[i].expected {
				fmt.Printf("Test case %d passed. Modified String: %v. Expected String: %v. \n", i+1, *testCases[i].messageIn, *testCases[i].expected)
			} else {
				fmt.Printf("Test case %d Failed. Modified String: %v. Expected String: %v. \n", i+1, *testCases[i].messageIn, *testCases[i].expected)
			}
		} else { // For when pointers are pointing to a null
			if testCases[i].messageIn == testCases[i].expected {
				fmt.Printf("Test case %d passed. Modified String: %v. Expected String: %v. \n", i+1, testCases[i].messageIn, testCases[i].expected)
			} else {
				fmt.Printf("Test case %d Failed. Modified String: %v. Expected String: %v. \n", i+1, testCases[i].messageIn, testCases[i].expected)
			}
		}
	}
}
