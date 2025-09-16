package main

import "fmt"

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func createUser(name string, membershipType string) User {
	var charLimit int
	if membershipType == "premium" {
		charLimit = 1000
	} else if membershipType == "standard" {
		charLimit = 100
	}

	// NOTE: Using "Membership: {" instead of "Membership: Membership{" on line 26 throws an error because GO needs explicit definition of type when initializing an embedded struct
	newUser := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: charLimit,
		},
	}

	return newUser
}

func main() {
	newUser := createUser("Max", "basic")

	fmt.Println(newUser)
}
