package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, singleMsg := range msg {
		for _, singleBadWord := range badWords {
			if singleMsg == singleBadWord {
				return i
			}
		}
	}

	return 0
}

func main() {
	messages := []string{"I ", "hate", "b*tch", "this", "b*tch"}
	badWords := []string{"b*tch"}

	badWordIndex := indexOfFirstBadWord(messages, badWords)
	fmt.Println(badWordIndex)
}
