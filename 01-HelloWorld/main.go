package main

import "fmt"

// init() is a special function in Go that runs before the main function
// It is typically used for initialization or setting up variables
func init() {
	fmt.Println("init")
}

// main is the main function in Go
// Every Go program must have a main function
// This is the entry point of the program execution
func main() {
	fmt.Println("Hello World")
}