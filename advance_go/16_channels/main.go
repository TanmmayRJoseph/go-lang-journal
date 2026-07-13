package main

import (
	"fmt"
	"time"
)

func main() {

	// ================================
	// CHANNELS (UNBUFFERED CHANNELS)
	// ================================

	// Till now we have learned:
	// 1. Goroutines -> run multiple tasks at the same time.
	// 2. WaitGroups -> wait until every goroutine finishes.
	//
	// Problem:
	// WaitGroups only wait.
	// They CANNOT send data back to the main function.
	//
	// So how do two goroutines communicate?
	//
	// Answer -> Channels.
	//
	// Think of a channel like a pipe.
	//
	//      Goroutine A  -------->  Channel  -------->  Goroutine B
	//
	// One goroutine puts data inside the pipe.
	// Another goroutine receives that data.
	//
	// This is called communication between goroutines.

	type User struct {
		ID   int
		Name string
	}

	// Create an UNBUFFERED channel.
	//
	// This channel can carry only User values.
	//
	// "make(chan User)" means:
	// "Create a communication pipe that transfers User structs."
	ch := make(chan User)

	// Start another goroutine.
	// This runs independently from the main function.
	go func() {

		// Pretend we are doing something slow,
		// like fetching a user from a database or API.
		time.Sleep(200 * time.Millisecond)

		// Send a User into the channel.
		//
		// "<-" means "send".
		//
		// IMPORTANT:
		// Since this is an UNBUFFERED channel,
		// this line will STOP here until someone receives it.
		//
		// So this goroutine cannot continue
		// until main() receives the value.
		ch <- User{
			ID:   90,
			Name: "Tanmmay",
		}

		// Once main() receives the value,
		// this goroutine becomes unblocked
		// and finishes.
	}()

	// Meanwhile the main goroutine keeps executing.
	fmt.Println("Main: waiting to receive user...")

	// Receive data from the channel.
	//
	// "<-" now means "receive".
	//
	// This line also blocks.
	// If no data has arrived yet,
	// main() patiently waits here.
	//
	// As soon as the worker goroutine sends a User,
	// it is copied into variable 'u'.
	u := <-ch

	// Now both goroutines are happy:
	//
	// Worker:
	//   "I finally found someone to receive my data."
	//
	// Main:
	//   "I finally got the data I was waiting for."
	//
	// This is called a HANDSHAKE.
	//
	// Sender waits for Receiver.
	// Receiver waits for Sender.
	// Neither continues until both meet.

	fmt.Println("Main: received user", u.ID, u.Name)
}
