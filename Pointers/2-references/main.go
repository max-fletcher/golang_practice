package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

type testCase struct {
	messageIn string
	expected  string
}

func main() {
	testCases := []testCase{
		{
			"English, motherfubber, do you speak it?",
			"English, mother****er, do you speak it?",
		},
		{
			"Oh man I've seen some crazy ass shiz in my time...",
			"Oh man I've seen some crazy ass **** in my time...",
		},
		{
			"Does he look like a witch?",
			"Does he look like a *****?",
		},
	}

	// NOTE: Remember to not use "for i, testCase := range testCases {" because "testCase" doesn't
	for i := range testCases {
		removeProfanity(&testCases[i].messageIn)

		if testCases[i].messageIn == testCases[i].expected {
			fmt.Printf("Test case %d passed. Modified String: %s. Expected String: %s. \n", i+1, testCases[i].messageIn, testCases[i].expected)
		} else {
			fmt.Printf("Test case %d Failed. Modified String: %s. Expected String: %s. \n", i+1, testCases[i].messageIn, testCases[i].expected)
		}
	}
}
