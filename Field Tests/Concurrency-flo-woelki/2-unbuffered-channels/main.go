package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// NOTE: This is an example of unbuffered channels i.e No limits are defined when the channel was made. It is synchronous by nature i.e the sender will block the execution until the receiver is ready to receive the data.

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

// accept an orderChan channel that receives [a pointer to an Order struct]/*Order and a WaitGroup. The arrow denotes that here, we will only receive *Order via this channel and not send
// anything(which is usually the default i.e send and receive if no arrows are used).
func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the WaitGroup

	fmt.Println("\n--- PROCESS ORDERS ---")
	for order := range orderChan { // This loop is special. It tells the go runtime that we are looping through the *Orders received via this orderChannel until the channel is closed
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
	wg.Add(2) //Add the number of goroutines that the WaitGroup will run

	orderChan := make(chan *Order) // Make a channel that sends and/or accepts an Order struct

	go func() { // We are running this anonymous func as a goroutine
		defer wg.Done() // the keyword "defer" is used to execute a line of code at the end of this function
		for _, order := range generateOrders(20) {
			orderChan <- order // Tells golang to send an Order into the orderChan channel. Remember that this is a blocking operation so until the sent data is received by something, execution will not move forward
		}

		// We need to close the channel after sending is done i.e after the loop terminates. Else a deadlock will occur due to a receiver trying to receive data, but there is no sender in place(in this case). The vice versa can happen as well.
		// The best practice is to close the channel on the sender side for unbuffered channels.
		close(orderChan)

		fmt.Println("Done with generating orders")
	}()

	fmt.Println("Start Processing Orders")

	go processOrders(orderChan, &wg)

	wg.Wait() // Works like await. Will make the go runtime wait until the processes above finish i.e the counter/WaitGroup becomes zero

	fmt.Println("End Processing Orders")
}
