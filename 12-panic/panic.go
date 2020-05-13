package main

import (
	"fmt"
	"sync"
	"time"
)

// A wait group for goroutines
var wg sync.WaitGroup

// A target function for the goroutine which takes time to complete
func say(s string) {

	// We provide a suitable recovery function in case of panics
	defer cleanup()

	// If we don't tell the wait group we're done we'll get a deadlock
	// We defer after cleanup to ensure we run first
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)

		// Introduce a panic
		if i == 2 {
			panic("Oopsie - i is 2")
		}

	}

}

// A recovery function - gets called after every say()
func cleanup() {

	// We avoid placing wait group code here because anyone could
	// call this function and we'd change goroutine counts

	// Check for a recover situation from a panic
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup: ", r)
	}

}

func main() {

	// Remember that main runs in its own thread
	// so without a wait group the main thread will
	// exit before the others have completed

	// Add a count to the wait group
	wg.Add(1)

	// Just add 'go' in front of function to make it a goroutine
	go say("Hey")

	wg.Add(1)
	go say("There")

	// Let the wait group work its way to the end
	wg.Wait()

}
