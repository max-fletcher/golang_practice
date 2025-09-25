package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line
func updateBalance(customer *customer, transaction transaction) error {
	if transaction.transactionType != "deposit" && transaction.transactionType != "withdrawal" {
		return errors.New("unknown transaction type")
	}
	if transaction.amount > customer.balance {
		return errors.New("insufficient funds")
	}

	if transaction.transactionType == "deposit" {
		customer.balance += transaction.amount
		return nil
	}
	if transaction.transactionType == "withdrawal" {
		customer.balance -= transaction.amount
		return nil
	}

	return nil
}

type testCase struct {
	name            string
	initialCustomer customer
	transaction     transaction
	expectedBalance float64
	expectError     bool
	errorMessage    string
}

func main() {
	testCases := []testCase{
		{
			name:            "Deposit operation",
			initialCustomer: customer{id: 1, balance: 100.0},
			transaction:     transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit},
			expectedBalance: 150.0,
			expectError:     false,
		},
		{
			name:            "Withdrawal operation",
			initialCustomer: customer{id: 2, balance: 200.0},
			transaction:     transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 100.0,
			expectError:     false,
		},
		{
			name:            "insufficient funds for withdrawal",
			initialCustomer: customer{id: 3, balance: 50.0},
			transaction:     transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal},
			expectedBalance: 50.0,
			expectError:     true,
			errorMessage:    "insufficient funds",
		},
		{
			name:            "unknown transaction type",
			initialCustomer: customer{id: 4, balance: 100.0},
			transaction:     transaction{customerID: 4, amount: 50.0, transactionType: "unknown"},
			expectedBalance: 100.0,
			expectError:     true,
			errorMessage:    "unknown transaction type",
		},
	}

	for i := range testCases {
		err := updateBalance(&testCases[i].initialCustomer, testCases[i].transaction)
		// If error occurs, check expected error and error string
		if err != nil && testCases[i].expectError && err.Error() == testCases[i].errorMessage && testCases[i].initialCustomer.balance == testCases[i].expectedBalance {
			fmt.Printf("Test case %d Passed.\n Customer after transaction: %v.\n Transaction: %v.\n Expected balance: %f.\n Expect error: %v.\n Error message: %s.\n Block 1.\n",
				i+1, testCases[i].initialCustomer, testCases[i].transaction, testCases[i].expectedBalance, testCases[i].expectError, testCases[i].errorMessage)
			// If no error occurs, check customer balance
		} else if err == nil && !testCases[i].expectError && testCases[i].transaction.transactionType == "deposit" && testCases[i].expectedBalance == (testCases[i].initialCustomer.balance) {
			fmt.Printf("Test case %d Passed.\n Customer after transaction: %v.\n Transaction: %v.\n Expected balance: %f.\n Expect error: %v.\n Error message: %s.\n Block 2.\n",
				i+1, testCases[i].initialCustomer, testCases[i].transaction, testCases[i].expectedBalance, testCases[i].expectError, testCases[i].errorMessage)
			// If no error occurs, check customer balance
		} else if err == nil && !testCases[i].expectError && testCases[i].transaction.transactionType == "withdrawal" && testCases[i].expectedBalance == (testCases[i].initialCustomer.balance) {
			fmt.Printf("Test case %d Passed.\n Customer after transaction: %v.\n Transaction: %v.\n Expected balance: %f.\n Expect error: %v.\n Error message: %s.\n Block 3.\n",
				i+1, testCases[i].initialCustomer, testCases[i].transaction, testCases[i].expectedBalance, testCases[i].expectError, testCases[i].errorMessage)
			// Else, you done messed up fam
		} else {
			fmt.Printf("Oopsie Doopsie !! Test case %d.\n Initial customer: %v.\n Transaction: %v.\n Expected balance: %f.\n Expect error: %v.\n Error message: %s.\n Block 4.\n",
				i+1, testCases[i].initialCustomer, testCases[i].transaction, testCases[i].expectedBalance, testCases[i].expectError, testCases[i].errorMessage)
		}
	}
}
