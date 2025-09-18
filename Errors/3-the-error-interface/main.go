package main

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %f by zero", de.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	val, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result: ", val)
}
