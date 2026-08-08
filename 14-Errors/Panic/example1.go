package main

import "fmt"

func main() {
	GetEmployeeInfo("", "", 0)
	CalculateTax(0)

}
func GetEmployeeInfo(firstName, lastName string, Salary int) float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Recovered from panic")
		}
	}()
	if firstName == "" {
		panic("firstName is empty")
	}
	if lastName == "" {
		panic("lastName is empty")
	}
	if Salary < 0 {
		panic("salary is less than zero")
	}
	return CalculateTax(Salary)
}
func CalculateTax(salary int) float64 {
	tax := float64(salary) * 0.09
	if tax > 1000 {
		panic("tax is greater than 1000")
	}
	return tax
}
