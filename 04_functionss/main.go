package main

import "fmt"

// add function
func add(x int, y int) int {
	return x + y
}

// multiple return values function
func sumAndProduct(x int, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}

// named return value function
func divide(x, y int) (add int, subtract int, multiply int, divide float64) {
	add = x + y
	subtract = x - y
	multiply = x * y

	if y != 0 {
		divide = float64(x) / float64(y)
	}

	return
}

// variadic function
func sumAll(numbers ...int) int {
	// used to take n number of arguments and return the sum of all numbers
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println("FUNCTIONS")

	sum1 := add(5, 10)
	fmt.Println("Sum of 5 and 10 is:", sum1)

	sum2, product := sumAndProduct(5, 10)
	fmt.Println("Sum of 5 and 10 is:", sum2)
	fmt.Println("Product of 5 and 10 is:", product)

	// using named return values
	addResult, subtractResult, multiplyResult, divideResult := divide(10, 5)
	fmt.Println("Add Result:", addResult)
	fmt.Println("Subtract Result:", subtractResult)
	fmt.Println("Multiply Result:", multiplyResult)
	fmt.Println("Divide Result:", divideResult)

	// varidic function
	numbers := []int{1, 2, 3, 4, 5}

	result := sumAll(numbers...)
	fmt.Println("Sum of all numbers is:", result)

	// anonymous function
	varFunc := func(x, y int) int {
		return x + y
	}

	result2 := varFunc(10, 20)
	fmt.Println("Result of anonymous function is:", result2)

}
