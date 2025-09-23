package main

import "fmt"

func isValidPassword(password string) bool {
	if len(password) <= 5 && len(password) >= 12 {
		return false
	}

	containsUppercase := false
	containsNum := false
	for _, char := range password {
		// Contains at least one uppercase letter
		if char >= 65 && char <= 90 {
			containsUppercase = true
		}
		// Contains at least one digit
		if char >= 48 && char <= 57 {
			containsNum = true
		}
	}

	return containsUppercase && containsNum
}

func main() {
	res := isValidPassword("Badmove123")
	fmt.Println(res)
}
