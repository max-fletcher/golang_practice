package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	contractor1 := contractor{
		name:         "Contractor boi",
		hourlyPay:    10,
		hoursPerYear: 20,
	}

	fmt.Println(contractor1.getName(), contractor1.getSalary())

	fullTime1 := fullTime{
		name:   "Fulltime boi",
		salary: 30,
	}

	fmt.Println(fullTime1.getName(), fullTime1.getSalary())
}
