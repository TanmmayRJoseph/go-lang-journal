package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {

		select {

		// If context is cancelled...
		case <-ctx.Done():

			// ctx.Err() tells us WHY it stopped.
			fmt.Println("Worker stopped:", ctx.Err())
			return

		// Otherwise continue working.
		default:

			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {

	// Create a root context.
	ctx := context.Background()

	// Create a cancellable context.
	ctx, cancel := context.WithCancel(ctx)

	// Start worker.
	go worker(ctx)

	// Let worker run for 3 seconds.
	time.Sleep(3 * time.Second)

	fmt.Println("Main: cancelling worker")

	// Tell the worker to stop.
	cancel()

	// Wait to see the output.
	time.Sleep(time.Second)
}