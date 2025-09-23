package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes") // Zero/empty map is nil
	}

	users := map[string]user{}
	for i, name := range names {
		users[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}

type test struct {
	names        []string
	phoneNumbers []int
}

func main() {
	case1 := test{
		names:        []string{"Eren", "Armin", "Mikasa"},
		phoneNumbers: []int{14355550987, 98765550987, 18265554567},
	}

	fmt.Println("Case1")
	fmt.Println(getUserMap(case1.names, case1.phoneNumbers))

	case2 := test{
		names:        []string{"Eren", "Armin"},
		phoneNumbers: []int{14355550987, 98765550987, 18265554567},
	}

	fmt.Println("Case2")
	fmt.Println(getUserMap(case2.names, case2.phoneNumbers))

	case3 := test{
		names:        []string{"George", "Annie", "Reiner", "Sasha"},
		phoneNumbers: []int{20955559812, 38385550982, 48265554567, 16045559873},
	}

	fmt.Println("Case3")
	fmt.Println(getUserMap(case3.names, case3.phoneNumbers))

	case4 := test{
		names:        []string{"George", "Annie", "Reiner"},
		phoneNumbers: []int{20955559812, 38385550982, 48265554567, 16045559873},
	}

	fmt.Println("Case4")
	fmt.Println(getUserMap(case4.names, case4.phoneNumbers))

	return
}
