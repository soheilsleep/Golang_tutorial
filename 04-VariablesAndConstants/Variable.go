package main
//method  1: Explicit type declaration
// This is the most explicit way to declare variables
// You specify both the type and the value
var i int = 12
var s string = "first name"

//method 2: Variable declaration without initialization
// The variable is declared but not initialized
// It will have the zero value for its type (0 for int, "" for string, etc.)
var i2 int
var s2 string

//method 3: Type inference
// Go automatically infers the type from the value
// Useful when you don't want to explicitly specify the type
var i3 = 22
var f1 = 12.66
var s3 = "hello world"

//method 4: Short variable declaration (:=)
// This is the most common and concise way to declare variables
// The := operator combines variable declaration and initialization
// It can only be used inside functions, not at package level
i4 := 12
f2 := 12.21
s4 := "Hello World"

//method 5: Group variables with multiple assignments
// You can declare and initialize multiple variables in one statement
// The := operator is used for short variable declaration
i5, f3, s5 = 25, 66.66, "Hello again"
// You can also reassign values to existing variables
i5, f3, s5 = 25, 66.66, "Hello again with out (var)"

// method 6: Group variables using parentheses
// This is useful for declaring multiple related variables
// Each variable can have its own type and value
var (
	i6 = 45
	f4 = 45.35
	s6 = "the end"
)