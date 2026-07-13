package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// waitgroups -> counter lets main wait to complete set of go routines
	// add how many routines u have

	var wg sync.WaitGroup

	// i will wait for 3 routines
	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("task 1")
		time.Sleep(250 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("task 2")
		time.Sleep(150 * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("task 3")
		time.Sleep(200 * time.Millisecond)
	}()

	fmt.Println("Main is now waiting for all tasks to finish")

	wg.Wait()

	fmt.Println("All tasks completed")
}
