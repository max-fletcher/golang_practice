package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}

// Methods need to defined outside of main()
// the first set of brackets after "func" is called a receiver and it signals to GO that we are appending a method to a struct(in this case, authenticationInfo)
// This is similar to classes in OOP having instance variables and methods
func (authInfo authenticationInfo) getAuthInfo() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

func main() {
	authUser := authenticationInfo{
		username: "Fletcher",
		password: "Password",
	}

	fmt.Println(authUser.getAuthInfo())
}
