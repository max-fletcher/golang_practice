package main

import (
	"fmt"
	"time"
)

// This is a normal function
func sendMessage(msg message) (string, int) {
	return msg.getMessage(), (len(msg.getMessage()) * 3)
}

// "message" interface
type message interface {
	getMessage() string
}

// "birthdayMessage" struct that contains "getMessage()" (see below)
type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

// "birthdayMessage" contains "getMessage()" so it is considered as implements "message" and can be considered as an instance of "message"
func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

// "sendingReport" struct that contains "getMessage()" (see below)
type sendingReport struct {
	reportName    string
	numberOfSends int
}

// "sendingReport" contains "getMessage()" so it is considered as implements "message" and can be considered as an instance of "message"
func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	bMsg := birthdayMessage{
		recipientName: "Dean Jones",
	}

	sentMsg1, sentLength1 := sendMessage(bMsg)
	fmt.Println(sentMsg1, sentLength1)

	sReport := sendingReport{
		reportName:    "Dean Jones BDay Sent",
		numberOfSends: 2,
	}

	sentMsg2, sentLength2 := sendMessage(sReport)
	fmt.Println(sentMsg2, sentLength2)
}
