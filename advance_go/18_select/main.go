package main

import (
	"fmt"
	"time"
)

func main() {

	// This channel will receive the result from our worker goroutine.
	resultChan := make(chan string)

	// Start a separate goroutine.
	go func() {

		// Pretend we are calling an external API.
		// This takes 400 milliseconds.
		time.Sleep(400 * time.Millisecond)

		// After the work is complete,
		// send the result into the channel.
		resultChan <- "worker: success"

	}()

	// Create a timeout channel.
	//
	// time.After() waits for the given duration
	// and then automatically sends the current time
	// into the returned channel.
	//
	// Nobody sends anything manually.
	timeoutCh := time.After(250 * time.Millisecond)

	// select waits until ONE of the channel operations
	// becomes ready.
	select {

	// If resultChan receives data first...
	case res := <-resultChan:
		fmt.Println("Received:", res)

	// If timeout channel receives first...
	case <-timeoutCh:
		fmt.Println("Timeout! Worker took too long.")
	}

	fmt.Println("Main function finished.")
}