package main

import "fmt"

// main is the entry point of the program
// This demonstrates different ways to print output in Go
func main() {
	// Define some variables of different types
	name := "soheil"
	age := 30
	nationalCode := 123456789
	score := 10.5
	
	// print: prints to stdout without a newline
	// Multiple arguments are automatically separated by spaces
	print("my name is ", name, " and age is ", age, " and national code is ", nationalCode, " and score is ", score, "\n")
	
	// println: prints to stdout with a newline
	// Multiple arguments are automatically separated by spaces
	println("my name is", name, "and age is", age, "and score is", score)
	
	// fmt.Printf: formatted print
	// Uses format specifiers to control output format
	// %s: string, %d: integer, %f: float
	fmt.Printf("my name is %s and age is %d and national code is %d and score is %f \n", name, age, nationalCode, score)
	
	// %T: prints the type of a variable
	fmt.Printf("name : my type is %T \n", name)
	
	// %b: prints an integer in binary format
	fmt.Printf("national code binary is : %b", nationalCode)

}