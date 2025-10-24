package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
		fmt.Printf("Received message from database %v\n", i+1)
	}
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}

func main() {
	count := 5
	ch, _ := getDBsChannel(count)
	waitForDBs(count, ch)

	// If you use these 2 lines, the main thread(and waitForDBs) might finish executing when count == 0 so it might not run properly(i.e Will run 0 times and not 5 times)
	// and might not show the fmts properly
	// ch, count := getDBsChannel(3)
	// waitForDBs(count, ch)
}
