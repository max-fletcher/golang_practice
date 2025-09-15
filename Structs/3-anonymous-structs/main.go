package main

import "fmt"

type car struct {
	brand   string
	model   string
	doors   int
	mileage int
	// wheel is a field containing an anonymous struct. It is not named and/or declared; just directly used.
	wheel struct {
		radius   int
		material string
	}
}

func main() {
	var myCar = car{
		brand:   "Rezvani",
		model:   "Vengeance",
		doors:   4,
		mileage: 35000,
		wheel: struct {
			radius   int
			material string
		}{
			radius:   35,
			material: "alloy",
		},
	}

	myCar.brand = "Volvo"
	myCar.brand = "GTP"
	myCar.doors = 4
	myCar.mileage = 10

	fmt.Printf(" Brand: %s\n model: %s\n doors: %d\nmileage: %d\n wheel-radius: %d\n wheel-material: %s\n", myCar.brand, myCar.model, myCar.doors, myCar.mileage, myCar.wheel.radius, myCar.wheel.material)
}
