package main // Every executable Go program starts with package main

import (
	"fmt"     // Used for printing messages and creating formatted errors
	"log"     // Used for logging errors and stopping the program
	"strconv" // Used for converting strings into numbers
)

// parseLevel converts a string into an integer
// and checks whether it is between 1 and 10.
func parseLevel(s string) (int, error) {

	// strconv.Atoi means "ASCII to Integer"
	// Example:
	// "5"  -> 5
	// "10" -> 10
	// "abc" -> ERROR

	n, err := strconv.Atoi(s)

	// If conversion failed...
	// err will contain information about what went wrong.
	if err != nil {

		// We return:
		// 0 as a dummy integer (because function must return an int)
		// and an error explaining the problem.
		return 0, fmt.Errorf("parseLevel: %v", err)
	}

	// Now we know the conversion succeeded.
	// Let's validate the value.

	// We only allow levels from 1 to 10.
	if n < 1 || n > 10 {

		// Create our own custom error.
		return 0, fmt.Errorf(
			"parseLevel: level must be between 1 and 10, got %d",
			n,
		)
	}

	// Everything is valid.
	// Return the number and nil error.
	return n, nil
}

// runLevel uses parseLevel before running the level.
func runLevel(s string) error {

	// Call parseLevel().
	// It returns:
	// input -> integer
	// err   -> possible error
	input, err := parseLevel(s)

	// If parseLevel failed,
	// stop this function immediately.
	if err != nil {

		// Wrap the original error with more context.
		return fmt.Errorf("runLevel: %v", err)
	}

	// If everything succeeded...
	fmt.Printf("Running level %d\n", input)

	// nil means "No error occurred."
	return nil
}

func main() {

	fmt.Println("ERROR HANDLING")

	// Call runLevel().
	// The := creates the variable err only inside this if statement.
	if err := runLevel("5"); err != nil {

		// log.Fatal()
		// 1. Prints the error
		// 2. Stops the program immediately
		log.Fatal("Error:", err)
	}
}
