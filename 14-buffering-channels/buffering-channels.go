package main

import (
	"fmt"
	"sync"
)

/*
Buffering and iterating over channels is a more real world use case.
*/

var wg sync.WaitGroup

// Send someValue to the channel after multiplying
func multiply(channel chan int, someValue int) {

	defer wg.Done()

	// Read this is "take the value someValue * 5 and send it to channel"
	channel <- someValue * 5

}

func main() {

	// Create a channel with a buffer for 10 items
	channel := make(chan int, 10)

	// Send 10 messages to the channel
	// No guarantee of order of execution here
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go multiply(channel, i)
	}

	// The waitgroup ensures the goroutines complete before moving on
	wg.Wait()
	close(channel)

	for item := range channel {
		fmt.Println(item)
	}
}
