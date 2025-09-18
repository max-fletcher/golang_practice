package main

import "fmt"

type cost struct {
	day   int
	value float64
}

// A function that is accepting a cost slice(containing slices) and a day int
func getDayCosts(costs []cost, day int) []float64 {
	foundCosts := []cost{} // A slice of structs(cost). Just for show
	totalCosts := []float64{}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			foundCosts = append(foundCosts, costs[i])
			totalCosts = append(totalCosts, costs[i].value)
		}
	}

	fmt.Println("foundCosts", foundCosts) // Print slice of cost structs

	return totalCosts
}

func main() {
	// A slice of structs
	costs := []cost{
		{
			day:   1,
			value: 10,
		},
		{
			day:   2,
			value: 20,
		},
		{
			day:   2,
			value: 30,
		},
		{
			day:   3,
			value: 10,
		},
		{
			day:   2,
			value: 30,
		},
	}

	totalCosts := getDayCosts(costs, 2)
	fmt.Println("totalCosts", totalCosts)
}
