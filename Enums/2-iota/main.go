package main

import (
	"fmt"
)

type emailStatus int

const (
	EmailBounced emailStatus = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)

func main() {
	type testCase struct {
		status   emailStatus
		expected string
	}

	runCases := []testCase{
		{EmailBounced, "EmailBounced"},
		{EmailInvalid, "EmailInvalid"},
		{EmailDelivered, "EmailDelivered"},
		{EmailOpened, "EmailOpened"},
		{17, "Unknown"},
	}

	testCases := runCases
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getEmailStatusName(test.status)
		if output != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
				Test Failed:
				status:   %v
				expected: %v
				actual:   %v
				`, test.status, test.expected, output,
			)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
				Test Passed:
				status:   %v
				expected: %v
				actual:   %v
				`, test.status, test.expected, output,
			)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func getEmailStatusName(status emailStatus) string {
	switch status {
	case EmailBounced:
		return "EmailBounced"
	case EmailInvalid:
		return "EmailInvalid"
	case EmailDelivered:
		return "EmailDelivered"
	case EmailOpened:
		return "EmailOpened"
	default:
		return "Unknown"
	}
}
