package main

import (
	"fmt"
)

type biller[C customer] interface {
	Name() string
	// Charge will accept either an "org" or a "user", both of which satisfy the "customer" interface. So we will be using the generic for customer
	Charge(C) bill
}

// don't edit below this line

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Name() string {
	return fmt.Sprintf("%s user biller", ub.Plan)
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

func main() {
	fmt.Println("Starting Org Biller Tests")

	type testCase1 struct {
		biller         orgBiller
		customer       org
		expectedAmount float64
		expectedEmail  string
	}

	runCases1 := []testCase1{
		{
			biller: orgBiller{Plan: "pro"},
			customer: org{
				Admin: user{UserEmail: "jaskier@oxenfurt.com"},
				Name:  "Oxenfurt",
			},
			expectedAmount: 3000,
			expectedEmail:  "jaskier@oxenfurt.com",
		},
		{
			biller: orgBiller{Plan: "basic"},
			customer: org{
				Admin: user{UserEmail: "vernon@temeria.com"},
				Name:  "Temeria",
			},
			expectedAmount: 2000,
			expectedEmail:  "vernon@temeria.com",
		},
		{
			biller: orgBiller{Plan: "pro"},
			customer: org{
				Admin: user{UserEmail: "fringilla@nilfgaard.com"},
				Name:  "Nilfgaard",
			},
			expectedAmount: 3000,
			expectedEmail:  "fringilla@nilfgaard.com",
		},
	}

	testCases1 := runCases1

	passCount1 := 0
	failCount1 := 0

	for i, test := range testCases1 {
		err := testBiller(test.biller, test.customer, test.expectedAmount, test.expectedEmail)
		if err != nil {
			failCount1++
			fmt.Printf(`--------------------------------- 
				OrgTest %d 
				Failed: %v`,
				i, err,
			)
		} else {
			passCount1++
			fmt.Printf(`---------------------------------
				OrgTest %d 
				Passed:
				biller:   %v
				customer: %v
				`, i, test.biller, test.customer,
			)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("OrgBilling: %d passed, %d failed \n\n", passCount1, failCount1)

	fmt.Println("Starting User Biller Tests")

	type testCase2 struct {
		biller         userBiller
		customer       user
		expectedAmount float64
		expectedEmail  string
	}

	runCases2 := []testCase2{
		{
			biller:         userBiller{Plan: "basic"},
			customer:       user{UserEmail: "vesemir@kaermorhen.com"},
			expectedAmount: 50,
			expectedEmail:  "vesemir@kaermorhen.com",
		},
		{
			biller:         userBiller{Plan: "pro"},
			customer:       user{UserEmail: "zoltan@mahakam.com"},
			expectedAmount: 100,
			expectedEmail:  "zoltan@mahakam.com",
		},
		{
			biller:         userBiller{Plan: "pro"},
			customer:       user{UserEmail: "extra@submit.com"},
			expectedAmount: 100,
			expectedEmail:  "extra@submit.com",
		},
	}

	testCases2 := runCases2
	passCount2 := 0
	failCount2 := 0

	for i, test := range testCases2 {
		err := testBiller(test.biller, test.customer, test.expectedAmount, test.expectedEmail)
		if err != nil {
			failCount2++
			fmt.Printf(`---------------------------------
				UserTest %d 
				Failed: %v
				`, i, err,
			)
		} else {
			passCount2++
			fmt.Printf(`---------------------------------
				UserTest %d 
				Passed:
				biller:   %v
				customer: %v
				`, i, test.biller, test.customer,
			)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("UserBilling: %d passed, %d failed\n", passCount2, failCount2)
}

func testBiller[C customer](
	b biller[C],
	c C,
	expectedAmount float64,
	expectedEmail string,
) error {
	currentBill := b.Charge(c)
	name := b.Name()

	if currentBill.Amount != expectedAmount ||
		currentBill.Customer.GetBillingEmail() != expectedEmail {
		return fmt.Errorf(`biller "%v" FAILED:
			biller Type:     %T
			customer Type:   %T
			customer:        %v
			expected amount: %v
			expected email:  %v
			actual amount:   %v
			actual email:    %v
			`,
			name,
			b,
			c,
			c,
			expectedAmount,
			expectedEmail,
			currentBill.Amount,
			currentBill.Customer.GetBillingEmail(),
		)
	}

	return nil
}
