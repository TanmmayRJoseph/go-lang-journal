package main

import (
	"errors"
	"fmt"
)

// This function pretends to do some important work.
// The 'success' variable decides whether the work succeeds or fails.
func doWork(success bool) error {

	// Imagine we opened a game, a file, or connected to a database.
	// Before using it, we first "acquire" (get) the resource.
	fmt.Println("start: resources acquired")

	// ------------------- DEFER -------------------
	// Think of defer like writing a reminder note.
	//
	// Example:
	// Your mom says:
	// "After you're done playing, put the toys back."
	//
	// Instead of remembering later, you immediately write a reminder.
	//
	// That's exactly what defer does.
	//
	// Go remembers this line and waits.
	// It DOES NOT run it now.
	// It only runs it right before this function ends.
	//
	// Even if an error happens...
	// Even if we return early...
	// This line WILL still run.
	defer fmt.Println("cleanup: resources released")

	// Pretend this is the important work.
	fmt.Println("work: doing something important")

	fmt.Println("work is done")

	// What if something goes wrong?
	if !success {

		// We return an error here...
		// BUT before leaving this function,
		// Go checks:
		//
		// "Did someone use defer?"
		//
		// Yes!
		//
		// So it runs the deferred line FIRST,
		// then returns the error.
		return errors.New("error: something went wrong")
	}

	// If everything went well,
	// the function is about to end.
	//
	// Again Go checks:
	// "Any deferred functions?"
	//
	// Yes!
	//
	// It prints the cleanup message,
	// then returns nil.
	return nil
}

func main() {

	fmt.Println("====== Understanding defer ======")

	fmt.Println("defer means:")
	fmt.Println("👉 'Run this later, just before this function finishes.'")

	// ---------------- SUCCESS CASE ----------------
	fmt.Println("\nCase 1: Everything works")

	if err := doWork(true); err != nil {
		fmt.Println("Error:", err)
	}

	// ---------------- ERROR CASE ----------------
	fmt.Println("\nCase 2: Something fails")

	if err := doWork(false); err != nil {
		fmt.Println("Error:", err)
	}
}
