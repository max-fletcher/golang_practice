package main

import (
	"fmt"
	"time"
)

// A rule of thumb is to NEVER run a function that accepts a channel arg in main(or maybe in the main execution thread). Just run such functions in a goroutine.
type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func main() {
	cases := [3][3]email{
		[3]email{
			{
				body: "Words are pale shadows of forgotten names. As names have power, words have power.",
				date: time.Date(2019, 2, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's like everyone tells a story about themselves inside their own head.",
				date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Bones mend. Regret stays with you forever.",
				date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		[3]email{
			{
				body: "Music is a proud, temperamental mistress.",
				date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "Have you heard of that website Boot.dev?",
				date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "It's awesome honestly.",
				date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
			},
		},
		[3]email{
			{
				body: "I have stolen princesses back from sleeping barrow kings.",
				date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I burned down the town of Trebon",
				date: time.Date(2019, 6, 6, 0, 0, 0, 0, time.UTC),
			},
			{
				body: "I have spent the night with Felurian and left with both my sanity and my life.",
				date: time.Date(2022, 7, 0, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	caseResults := [3][3]bool{
		[3]bool{true, false, false},
		[3]bool{true, true, true},
		[3]bool{true, true, false},
	}

	for i := 0; i < len(cases); i++ {
		result := checkEmailAge(cases[i])

		if result == caseResults[i] {
			fmt.Printf("Test %d passed.\n", i)
		}
	}
}
