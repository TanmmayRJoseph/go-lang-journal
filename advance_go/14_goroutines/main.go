package main

import (
	"fmt"
	"time"
)

func main() {
	// concepts = doing multiple tasks in a way that they can overlap in time
	// overlap = while task A (network ,db) Task B can run

	// overlapping -> concurrency
	// executing multiple task at the exact same time

	// backend -> multiple service calls -> combine all results

	// go routine
	// concurrent execution of a function

	start := time.Now()

	// GO a
	go func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("go routine", time.Since(start))
	}()

	// GO b
	go func() {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("go routine", time.Since(start))
	}()

	fmt.Println("Main :started two go-routine", time.Since(start))

	fmt.Println("main doing step 1", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main doing step 2", time.Since(start))
	time.Sleep(100 * time.Millisecond)

	fmt.Println("main doing step 3", time.Since(start))

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main executed")



}
