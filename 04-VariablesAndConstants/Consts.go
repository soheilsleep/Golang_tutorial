package main

import (
	"fmt"
)

// main is the entry point of the program
// This demonstrates constants in Go
func main() {
	// Constants are values that cannot be changed after they are defined
	// They are declared using the 'const' keyword
	// Constants can be grouped together using parentheses
	const (
		name   = "soheilsleep"
		number = 26
		city   = "ahvaz"
		pi     = 3.14
	)

	// Print the constant values
	fmt.Printf("name : %s number : %d city:%s pi: %f \n", name, number, city, pi)

	// Constants can also be declared individually
	// This is useful for defining URLs, API endpoints, or other fixed values
	const GoogleBaseUrl = "https://www.google.com"
	const MapUrl = "/maps"

	// Print the individual constants
	fmt.Println(GoogleBaseUrl, MapUrl)

}