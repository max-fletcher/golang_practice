package main

import (
	"fmt"
	"regexp"
	"strings"
)

func cleanupText(text string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+") // Matches any character NOT a-z or A-Z
	cleanedText := reg.ReplaceAllString(text, "")
	return cleanedText
}

func countDistinctWords(messages []string) int {
	distinctWordsMap := make(map[string]struct{})
	for _, message := range messages {
		words := strings.Split(message, " ")

		for _, word := range words {
			cleanedWord := strings.ToLower(cleanupText(word))
			// fmt.Println("Word of the day:", cleanedWord)

			if cleanedWord != "" {
				_, exists := distinctWordsMap[cleanedWord]
				if !exists {
					distinctWordsMap[cleanedWord] = struct{}{}
				}
			}
		}
	}

	return len(distinctWordsMap)
}

type testCase struct {
	messages []string
	expected int
}

func main() {
	testCases := []testCase{
		{
			[]string{"WTS Arcanite Bar! Cheaper than AH", "Do you need an Arcanite Bar!"},
			10,
		},
		{
			[]string{"Could you give me a number crunch real quick?", "Looks like we have a 32.33% (repeating of course) percentage of survival."},
			19,
		},
		{
			[]string{"LFG UBRS", "lfg ubrs", "LFG Ubrs"},
			2,
		},
		{
			[]string{"Alright time's up! Let's do this.", "Leroy Jenkins!", "Damn it Leroy"},
			10,
		},
		{
			[]string{"I'm out of range", "I'm out of mana"},
			5,
		},
		{
			[]string{
				"LF9M UBRS need all",
				"LF8M UBRS need all",
				"LF7M UBRS need all",
				"LF6M UBRS need tanks and heals",
				"LF5M UBRS need tanks and heals",
				"LF4M UBRS need tanks and heals",
				"LF3M UBRS need tanks and healer",
				"LF2M UBRS need tanks",
				"LF1M UBRS need tank",
				"Group is full thanks!",
			},
			21,
		},
		{
			[]string{""},
			0,
		},
	}

	for i, testCase := range testCases {
		distinctWordCount := countDistinctWords(testCase.messages)
		if distinctWordCount == testCase.expected {
			fmt.Printf("Test case %d: Passed | Counted: %d | Expected: %d \n", i+1, distinctWordCount, testCase.expected)
		} else {
			fmt.Printf("Test case %d: Failed | Counted: %d | Expected: %d \n", i+1, distinctWordCount, testCase.expected)
		}
	}
}
