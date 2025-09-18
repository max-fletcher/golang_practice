package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func main() {
	res, err := divide(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
