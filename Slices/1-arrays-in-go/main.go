package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgArr := [3]string{primary, secondary, tertiary}
	costArr := [3]int{len(primary), len(secondary), len(tertiary)}

	return msgArr, costArr
}

func main() {
	msgArr, costArr := getMessageWithRetries("Hello", "World", "Boi")
	fmt.Println(msgArr, costArr)
}
