package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS <KEY:VALUE>")
	fmt.Println("Syntax :- map[keyType]valueType ")

	// map[keyType]valueType

	// Declaring a map
	ages := map[string]int{
		"John":  25,
		"Jane":  30,
		"Bob":   35,
		"Alice": 28,
	}

	fmt.Println(ages)
	fmt.Println(ages["John"], len(ages))

	// creating a empty map using make
	var emptyMap = make(map[string]int) //nil map
	emptyMap["John"] = 25
	emptyMap["Jane"] = 30
	emptyMap["Bob"] = 35
	emptyMap["Alice"] = 28

	users := map[string]string{
		"John":  "Doe",
		"Jane":  "Smith",
		"Bob":   "Johnson",
		"Alice": "Williams",
	}

	// deleting a key-value pair from a map
	delete(users, "Bob")
	fmt.Println("map after deleting bob", users)

	fmt.Println()

	fmt.Println("value,ok pattern ")
	fmt.Println("standard convention in Go used to check the success of an operation that might fail or return a 0 value")

	points := map[string]int{
		"John":  10,
		"Jane":  20,
		"Bob":   30,
		"Alice": 40,
	}

	// value, ok pattern to check if a key exists in the map
	value, ok := points["John"]
	if ok {
		fmt.Println("Value found:", value)
	} else {
		fmt.Println("Value not found")
	}

}
