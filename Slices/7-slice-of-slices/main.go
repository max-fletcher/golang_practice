package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	// Using this line instead of "matrix := [][]int{}" is better since the make function makes a slice of fixed length, otherwise,
	// append will reallocate memory which impacts performance
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func main() {
	table1 := createMatrix(5, 10)
	fmt.Println(table1)

	table2 := createMatrix(3, 4)
	fmt.Println(table2)
}
