package main

import "fmt"

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage(analytics *Analytics, message Message) {
	analytics.MessagesTotal++
	if message.Success {
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesFailed++
	}
}

type testCase struct {
	initialAnalytics Analytics
	newMessage       Message
	expected         Analytics
}

func main() {
	testCases := []testCase{
		{
			initialAnalytics: Analytics{MessagesTotal: 0, MessagesFailed: 0, MessagesSucceeded: 0},
			newMessage:       Message{Recipient: "mickey", Success: true},
			expected:         Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 1, MessagesFailed: 0, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "minnie", Success: false},
			expected:         Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
		},
		{
			initialAnalytics: Analytics{MessagesTotal: 2, MessagesFailed: 1, MessagesSucceeded: 1},
			newMessage:       Message{Recipient: "goofy", Success: false},
			expected:         Analytics{MessagesTotal: 3, MessagesFailed: 2, MessagesSucceeded: 1},
		},
	}

	for i := range testCases {
		analyzeMessage(&testCases[i].initialAnalytics, testCases[i].newMessage)
		if testCases[i].initialAnalytics == testCases[i].expected {
			fmt.Printf("Test case %d passed. Modified struct: %v. Expected struct: %v. \n", i+1, testCases[i].initialAnalytics, testCases[i].expected)
		} else {
			fmt.Printf("Test case %d failed. Modified struct: %v. Expected struct: %v. \n", i+1, testCases[i].initialAnalytics, testCases[i].expected)
		}
	}
}
