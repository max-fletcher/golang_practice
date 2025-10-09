package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Start: Defining enums
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusShipped   OrderStatus = "shipped"
	StatusDelivered OrderStatus = "delivered"
	StatusCancelled OrderStatus = "cancelled"
)

// End: Defining enums

type Order struct {
	ID     int
	Status OrderStatus
	mu     sync.Mutex
}

func generateOrders(count int) []*Order { // returns an array of pointers to Order structs
	// # NOTE: "orders := make([]*Order, count)" makes an array of pointers that can hold "Order" slices; initial value of orders(fmt.Println(orders)) is [<nil>, <nil>, <nil>].
	// However, "orders := make([]Order, count)" makes an array of "Order" slices; initial value of orders(fmt.Println(orders)) is [{0 } {0 } {0 }] i.e a 0 for ID and empty string for Status.
	// You will have to set a default yourself if you want a default enum value i.e what we are doing here which is a function to create an array of pointers to Order structs with initial values
	orders := make([]*Order, count) // initialize an array of order pointer(structs)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ // Create and add reference to the orders pointer slice
			ID:     i + 1,
			Status: StatusPending,
		}
	}

	return orders
}

func processOrders(orders []*Order) { // accept a array of pointers to Order structs
	fmt.Println("\n--- PROCESS ORDERS ---")
	for _, order := range orders {
		// Simulating some work being done using the sleep function
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Process order no. %d.\n", order.ID)
	}
	fmt.Println("-----------------------------")
}

// NOTE: This is replaced with the function below so that we can test goroutines more effectively
func updateOrderStatuses(orders []*Order) { // accept ar array of pointers to Order structs
	fmt.Println("\n--- UPDATE ORDERS STATUS ---")
	for _, order := range orders {
		// Simulating some work being done using the sleep function
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		// status := []OrderStatus{StatusPending, StatusShipped, StatusDelivered, StatusCancelled}[rand.Intn(4)]
		// # NOTE: The one line above can replace the 3 lines of code below
		// START: Generate random order status
		statuses := []OrderStatus{StatusPending, StatusShipped, StatusDelivered, StatusCancelled}
		randomIndex := rand.Intn(4)
		status := statuses[randomIndex]
		// END: Generate random order status
		order.Status = status // update order status

		fmt.Printf("Updated status for order no. %d, new status: %s.\n", order.ID, order.Status)
	}
	fmt.Println("-----------------------------")
}

func reportOrderStatus(orders []*Order) { // accept ar array of pointers to Order structs
	fmt.Println("\n--- REPORT STATUSES ---")
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")
		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}
		fmt.Println("-----------------------------")
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3) //Add the number of goroutines that the WaitGroup will run

	orders := generateOrders(20)
	fmt.Println("Orders after initialization", orders)

	fmt.Println("Start Processing Orders")

	go func() { // We are running this anonymous func as a goroutine
		defer wg.Done() // the keyword "defer" is used to execute a line of code at the end of this function
		processOrders(orders)
	}()

	// go func() { // We are running this anonymous func as a goroutine
	// 	defer wg.Done() // the keyword "defer" is used to execute a line of code at the end of this function
	// 	updateOrderStatuses(orders)
	// }()

	// NOTE: We are using the above anonymous functions instead of the lines below because otherwise, we won't be able to see the fmt.Print statements in the terminal
	// go processOrders(orders)
	// go updateOrderStatuses(orders)
	// go reportOrderStatus(orders)

	wg.Wait() // Works like await. Will make the go runtime wait until the processes above finish i.e the counter/WaitGroup becomes zero

	reportOrderStatus(orders)
	fmt.Println("End Processing Orders")
}
