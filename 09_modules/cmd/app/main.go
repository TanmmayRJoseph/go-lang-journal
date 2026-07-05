package main

import (
	"fmt"
	"modules/internal/greet"
)

func main() {
	msg1 := greet.Greet("Tanmay")
	fmt.Println(msg1)
}
