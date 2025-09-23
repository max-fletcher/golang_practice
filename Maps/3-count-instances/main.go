package main

import (
	"fmt"
)

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, msgUser := range messagedUsers {
		_, exists := validUsers[msgUser]
		if exists {
			validUsers[msgUser]++
		}
	}
}

func main() {
	testCase1 := []string{}
	validUsers1 := map[string]int{"tyrion": 0}
	updateCounts(testCase1, validUsers1)
	fmt.Println(validUsers1)

	testCase2 := []string{"cersei", "jaime", "tyrion"}
	validUsers2 := map[string]int{"tywin": 0}
	updateCounts(testCase2, validUsers2)
	fmt.Println(validUsers2)

	testCase3 := []string{"cersei", "cersei", "cersei", "tyrion"}
	validUsers3 := map[string]int{"cersei": 0}
	updateCounts(testCase3, validUsers3)
	fmt.Println(validUsers3)

	testCase4 := []string{"cersei", "tywin", "jaime", "cersei", "tyrion", "cersei", "jaime"}
	validUsers4 := map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0}
	updateCounts(testCase4, validUsers4)
	fmt.Println(validUsers4)

	testCase5 := []string{"cersei", "cersei", "jaime", "jaime", "tywin", "cersei", "tywin", "tyrion"}
	validUsers5 := map[string]int{"cersei": 0, "jaime": 0, "tyrion": 0}
	updateCounts(testCase5, validUsers5)
	fmt.Println(validUsers5)
}
