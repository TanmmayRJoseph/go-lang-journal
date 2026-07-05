package main

import "fmt"

func main() {
	fmt.Println("ARRAYS AND SLICES")
	fmt.Println("Arrays are fixed size in golang")
	fmt.Println("So in go lang we use slices")

	// declaring an array
	var arr [5]int

	arr[0] = 90
	arr[1] = 80
	arr[2] = 70
	arr[3] = 60
	arr[4] = 50

	fmt.Println(arr)

	// array literal
	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	fmt.Println()

	fmt.Println("Slices are dynamic in size")
	fmt.Println("Slices are most common collection types")

	// declaring a slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	// declaring a empty int slice
	var nums []int
	fmt.Println(nums)

	// appending values to a slice
	nums = append(nums, 10)
	nums = append(nums, 20)
	nums = append(nums, 30)
	nums = append(nums, 40)
	nums = append(nums, 50)

	// getting the length of a slice
	fmt.Println("Length of nums slice is:", len(nums))

	fmt.Println()

	fmt.Println("Capacity is how many elements you can store")

	scores := make([]int, 5, 10)
	fmt.Println("Length of scores slice is:", len(scores))   // length is 5
	fmt.Println("Capacity of scores slice is:", cap(scores)) // capacity is 10

	// slice concatenation
	slice1 := []int{10, 20, 30}
	slice2 := []int{40, 50, 60}

	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)


	// range on slices
	for index, value := range slice3 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	

}
