package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

func main() {
	type testCase struct {
		email string
		count int
	}

	runCases := []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
	}

	testCases := runCases

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		if output := sc.val(test.email); output != test.count {
			failCount++
			fmt.Printf(`---------------------------------
				Test Failed:
				email: %v
				count: %v
				expected count: %v
				actual count:   %v`,
				test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
				Test Passed:
				email: %v
				count: %v
				expected count: %v
				actual count:   %v
				`, test.email, test.count, test.count, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
