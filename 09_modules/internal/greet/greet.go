package greet

import (
	"fmt"
	"strings"
)

func printName(name string) string {
	clean := Greet(name)
	return fmt.Sprintf("Hello %s", clean)
}

// a function that prints the greeting message
func Greet(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "Guest"
	}

	return strings.ToUpper(name)
}
