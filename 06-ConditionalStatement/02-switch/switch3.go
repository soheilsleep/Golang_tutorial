package main

// the same if else file but with switch and case
import "fmt"

func main() {
	var Salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0
	fmt.Print("enter your salary: ")
	fmt.Scanln(&Salary)
	switch {
	case Salary <= minSalary:
		taxPercent = 0
	case Salary > minSalary && Salary <= minSalary*2:
		taxPercent = 0.5
	case Salary > minSalary*2 && Salary <= minSalary*3:
		taxPercent = 0.7
	case Salary > minSalary*3 && Salary <= minSalary*4:
		taxPercent = 0.1
	default:
		taxPercent = 0.15

	}
	fmt.Printf("your taxPercent is: %.2f\n", taxPercent*Salary)
	fmt.Printf("your salary is: %.2f\n ", Salary-taxPercent*Salary)
}
