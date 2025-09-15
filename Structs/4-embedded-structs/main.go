package main

import "fmt"

type sender struct {
	rateLimit int
}

// embedded struct i.e the sender has no name. It is directly accessible i.e
// newUserZ := user2{}
// newUserZ.rateLimit = 1000
type user1 struct {
	name   string
	number int
	sender
}

// Nested struct i.e the sender has a name declared so you need to pass through to access the nested properties i.e
// newUserZ := user2{}
// newUserZ.sender.rateLimit = 1000
type user2 struct {
	name   string
	number int
	sender sender
}

func main() {
	// Create an instance(newUser1) of an embedded struct
	newUser1 := user1{}
	newUser1.name = "Max"
	newUser1.number = 10
	newUser1.rateLimit = 200

	fmt.Printf(" Name: %s\n number: %d\n rateLimit: %d\n\n", newUser1.name, newUser1.number, newUser1.rateLimit)

	// ALSO an instance(shorthand) that creates an instance of embedded struct. Notice the comma at the end of "}" (line  43). It is needed to have this trailing comma else GO throws an error
	newUser2 := user1{
		name:   "Fletcher",
		number: 50,
		sender: sender{
			rateLimit: 300,
		},
	}

	fmt.Printf(" Name: %s\n number: %d\n rateLimit: %d\n\n", newUser2.name, newUser2.number, newUser2.sender.rateLimit)

	// An instance that uses nested struct
	newUser3 := user2{}
	newUser3.name = "Boi"
	newUser3.number = 70
	newUser3.sender.rateLimit = 500

	fmt.Printf(" Name: %s\n number: %d\n rateLimit: %d\n", newUser3.name, newUser3.number, newUser3.sender.rateLimit)
}
