package main

import (
	"fmt"
)

/*
Channels provide a blocking mechanism to send messages between concurrent goroutines.
*/

// Send someValue to the channel after multiplying
func multiply(channel chan int, someValue int) {

	// Read this is "take the value someValue * 5 and send it to channel"
	channel <- someValue * 5

}

func main() {

	// Make a channel
	channel := make(chan int)
	defer close(channel)

	// Create a couple of goroutines
	go multiply(channel, 5)
	go multiply(channel, 3)

	// Pop out the values from the channels
	// Note that we don't use any blocking
	// because the channel does it for us
	v1, v2 := <-channel, <-channel

	fmt.Println(v1, v2)
}
