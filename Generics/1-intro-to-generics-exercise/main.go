package main

import "fmt"

func getLast[T any](s []T) T {
	length := len(s)
	if length == 0 {
		return *new(T) // This is how you return a zero value for a generic of type T
	}

	return s[length-1]
}

func main() {
	emails := []string{"mail1@mail.com", "mail2@mail.com", "mail3@mail.com", "mail4@mail.com", "mail5@mail.com"}
	fmt.Println(getLast(emails))

	type testCase struct {
		input    interface{}
		expected interface{}
	}

	runCases := []testCase{
		{[]int{}, 0},
		{[]bool{true, false, true, true, false}, false},
	}

	testCases := append(runCases, []testCase{
		{[]int{}, 0},
		{[]bool{true, false, true, true, false}, false},
		{[]int{1, 2, 3, 4}, 4},
		{[]string{"a", "b", "c", "d"}, "d"},
	}...)

	passed, failed := 0, 0

	for _, test := range testCases {
		switch v := test.input.(type) {
		case []int:
			if output := getLast(v); output != test.expected {
				fmt.Errorf(`
					---------------------------------
					Test Failed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`
					---------------------------------
					Test Passed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []string:
			if output := getLast(v); output != test.expected {
				fmt.Errorf(`---------------------------------
					Test Failed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
					Test Passed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		case []bool:
			if output := getLast(v); output != test.expected {
				fmt.Errorf(`---------------------------------
					Test Failed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				failed++
			} else {
				fmt.Printf(`---------------------------------
					Test Passed:
					input:    %v
					expected: %v
					actual:   %v
					`,
					v,
					test.expected,
					output,
				)
				passed++
			}
		}
	}

	fmt.Println("---------------------------------")

	fmt.Printf("%d passed, %d failed", passed, failed)
}
