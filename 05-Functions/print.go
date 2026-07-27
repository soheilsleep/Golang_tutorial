package main

import "fmt"

func main() {
	name := "soheil"
	age := 30
	nationalCode := 123456789
	score := 10.5
	print("my name is ", name, " and age is ", age, " and national code is ", nationalCode, " and score is ", score, "\n")
	println("my name is", name, "and age is", age, "and score is", score)
	fmt.Printf("my name is %s and age is %d and national code is %d and score is %f \n", name, age, nationalCode, score)
	fmt.Printf("name : my type is %T \n", name)
	fmt.Printf("national code binary is : %b", nationalCode)

}
