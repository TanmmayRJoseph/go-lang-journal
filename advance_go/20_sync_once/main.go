package main

import (
	"fmt"
	"sync"
	"time"
)

/////////////////////////////////////////////////////////////////
// Shared Data
/////////////////////////////////////////////////////////////////

// Total processed orders.
// Multiple goroutines will modify this.
var totalOrders int

// Mutex protects writes to totalOrders.
var mu sync.Mutex

// RWMutex protects reading shared configuration.
var configMu sync.RWMutex

// Configuration shared by all workers.
var config = map[string]string{
	"DB": "Not Initialized",
}

// Ensures initialization happens only once.
var once sync.Once

/////////////////////////////////////////////////////////////////
// Initialization
/////////////////////////////////////////////////////////////////

func initializeDatabase() {

	fmt.Println("Initializing Database Connection...")

	// Simulate expensive initialization.
	time.Sleep(2 * time.Second)

	// Lock because we're writing shared config.
	configMu.Lock()
	config["DB"] = "Connected"
	configMu.Unlock()

	fmt.Println("Database Connected!")
}

/////////////////////////////////////////////////////////////////
// Worker
/////////////////////////////////////////////////////////////////

func processOrder(orderID int, wg *sync.WaitGroup) {

	// Tell WaitGroup we're finished
	// no matter how the function exits.
	defer wg.Done()

	// Even if 100 workers call this,
	// initializeDatabase() executes ONLY ONCE.
	once.Do(initializeDatabase)

	/////////////////////////////////////////////////////////////
	// Read shared configuration
	/////////////////////////////////////////////////////////////

	// Multiple workers can read simultaneously.
	configMu.RLock()

	fmt.Printf("Worker %d using DB Status: %s\n",
		orderID,
		config["DB"],
	)

	configMu.RUnlock()

	/////////////////////////////////////////////////////////////
	// Simulate processing
	/////////////////////////////////////////////////////////////

	fmt.Println("Processing Order:", orderID)

	time.Sleep(1 * time.Second)

	/////////////////////////////////////////////////////////////
	// Update shared variable
	/////////////////////////////////////////////////////////////

	// Only ONE goroutine can update totalOrders
	// at a time.
	mu.Lock()

	totalOrders++

	mu.Unlock()

	fmt.Println("Finished Order:", orderID)
}

/////////////////////////////////////////////////////////////////
// Main
/////////////////////////////////////////////////////////////////

func main() {

	// WaitGroup waits for ALL workers.
	var wg sync.WaitGroup

	// Launch 5 workers.
	for i := 1; i <= 5; i++ {

		// Increase WaitGroup counter.
		wg.Add(1)

		go processOrder(i, &wg)
	}

	fmt.Println("Main waiting for workers...")

	// Block until all workers call Done().
	wg.Wait()

	fmt.Println()

	fmt.Println("================================")
	fmt.Println("All Workers Finished")
	fmt.Println("Total Orders:", totalOrders)
	fmt.Println("================================")
}
