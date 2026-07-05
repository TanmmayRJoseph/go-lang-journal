package main

import "fmt"

func main() {
	fmt.Println("POINTERS")

	// score is a normal variable.
	// Imagine this is a box containing the number 10.
	score := 10

	fmt.Println("before:", score)

	// &score means:
	// "Take the ADDRESS of score."
	// Instead of sending the value (10),
	// we are sending where score lives in memory.
	addScore(&score)

	// Since addScore changed the ORIGINAL score,
	// the value here is also updated.
	fmt.Println("after:", score)
}

// score *int means:
// "I am expecting the ADDRESS of an integer."
func addScore(score *int) {

	// score is NOT the number itself.
	// It is just the address.

	// *score means:
	// "Go to that address and get the actual value."

	// Since we're changing the value at that address,
	// the original variable also changes.
	*score += 10
}
