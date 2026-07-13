package main

import (
	"fmt"
	"time"
)

// Rule to remember
// Unbuffered channel: "I can't send until someone receives."
// Buffered channel: "I can send until the buffer is full."
func main() {

	// Create a buffered channel.
	// It can store ONLY 2 strings before the sender has to wait.
	jobs := make(chan string, 2)

	// ---------------- Producer ----------------
	go func() {

		fmt.Println("Producer: Sending Job 1")

		// Buffer now contains:
		// [job1]
		// Producer DOES NOT wait because there is space.
		jobs <- "Job 1"

		fmt.Println("Producer: Sending Job 2")

		// Buffer now contains:
		// [job1][job2]
		// Buffer is now FULL.
		jobs <- "Job 2"

		fmt.Println("Producer: Sending Job 3")

		// There is NO SPACE left.
		// Producer will STOP here until consumer removes one item.
		jobs <- "Job 3"

		// Once consumer removes Job1,
		// Job3 gets stored automatically.
		fmt.Println("Producer: Finished Sending")

		// Close channel because no more jobs are coming.
		close(jobs)

	}()

	// ---------------- Consumer ----------------
	for job := range jobs {

		fmt.Println("Consumer received:", job)

		// Simulate slow work.
		time.Sleep(500 * time.Millisecond)

		fmt.Println("Consumer finished:", job)
	}

	fmt.Println("All jobs completed.")
}
