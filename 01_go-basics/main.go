package main

import (
	"fmt"
	"math"
)

// MAIN FUNCTION
func main() {

	fmt.Println("Here we go with GO again")

	var channelName string
	channelName = "TRJ Dev"
	fmt.Println("Channel Name", channelName)

	var year int = 2026
	fmt.Println("Year is ", year)

	// short hand syntax
	age := 50
	fmt.Println("Age is ", age)

	// packages example math package
	root := 5
	fmt.Println("Square root of ", root, " is ", math.Sqrt(float64(root)))

	// multiple variable declaration
	name, city, age := "TRJ", "Bangalore", 22
	fmt.Println("Name is ", name, " City is ", city, " Age is ", age)

	// constants used to declare values that cannot be changed
	const pi float64 = 3.14
	fmt.Println("Value of pi is ", pi)
	// groupping constants
	const (
		pi1 float64 = 3.14
		e   float64 = 2.71
	)
	fmt.Println("Value of pi is ", pi1, " Value of e is ", e)

	fmt.Println()

	fmt.Println("IF ELSE")
	score := 50
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}

	fmt.Println()

	fmt.Println("FOR LOOP")

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is ", i)
	}

	// example of while loop in go
	j := 0
	for j < 5 {
		fmt.Println("Value of j is ", j)
		j++
	}

	fmt.Println()

	fmt.Println("SWITCH CASE")
	// example of switch case in go
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Today is not Monday, Tuesday or Wednesday")
	}
}
