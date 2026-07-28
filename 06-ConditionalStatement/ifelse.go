package main

import "fmt"

func main() {
	var Salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0
	fmt.Print("enter your salary: ")
	fmt.Scanln(&Salary)
	if Salary <= minSalary {
		taxPercent = 0
	} else if Salary > minSalary && Salary <= minSalary*2 {
		taxPercent = 0.05
	} else if Salary > minSalary*2 && Salary <= minSalary*3 {
		taxPercent = 0.07
	} else if Salary > minSalary*3 && Salary <= minSalary*4 {
		taxPercent = 0.1
	} else {
		taxPercent = 0.15
	}
	fmt.Printf("your taxPercent is: %.2f\n", taxPercent*Salary)
	fmt.Printf("your salary is: %.2f\n ", Salary-taxPercent*Salary)
}
