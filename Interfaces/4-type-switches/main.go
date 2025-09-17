package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch val := e.(type) {
	case email:
		return val.toAddress, val.cost()
	case sms:
		return val.toPhoneNumber, val.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	email1 := email{
		isSubscribed: true,
		body:         "New Message lolololo",
		toAddress:    "New Address lolololo",
	}

	fmt.Println(getExpenseReport(email1))

	sms1 := sms{
		isSubscribed:  true,
		body:          "New Message lolololo",
		toPhoneNumber: "+293084293094",
	}

	fmt.Println(getExpenseReport(sms1))
}
