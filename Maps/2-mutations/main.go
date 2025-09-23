package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	_, exists := users[name] // Can also be written as "elem, ok := users[name]""

	if exists && users[name].scheduledForDeletion {
		delete(users, name)
		fmt.Println("After Delete", users)
		return true, nil
	}

	if exists && !users[name].scheduledForDeletion {
		return false, nil
	}

	return false, errors.New("not found")
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	case1 := map[string]user{"Erwin": {"Erwin", 14355550987, true}, "Levi": {"Levi", 98765550987, true}, "Hanji": {"Hanji", 18265554567, true}}
	delete1 := "Erwin"

	res1, error1 := deleteIfNecessary(case1, delete1)
	fmt.Println(res1, error1)

	case2 := map[string]user{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}}
	delete2 := "Erwin"

	res2, error2 := deleteIfNecessary(case2, delete2)
	fmt.Println(res2, error2)

	case3 := map[string]user{"Erwin": {"Erwin", 14355550987, false}, "Levi": {"Levi", 98765550987, false}, "Hanji": {"Hanji", 18265554567, false}}
	delete3 := "Eren"

	res3, error3 := deleteIfNecessary(case3, delete3)
	fmt.Println(res3, error3)

	return
}
