package main

import "fmt"

const (
	BaseSalary       = 6600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate  = 80000
	TaxRate          = 0.09
)

func main() {
	FullTimeEmployees := []FullTimeEmployee{
		{1, "soheilsleep", "1234567890", 40},
		{2, "taha dehnavi", "1548628460", 45},
	}
	PartTimeEmployee := []PartTimeEmployee{
		{3, "reza mohamadi", "1985181580", 160},
		{4, "ali gorbay", "4518456150", 120},
	}
	ShiftEmployee := []ShiftEmployee{
		{5, "hediye mohamadi", "7456274596", 50},
		{6, "asal gorbay", "1462984573", 60},
	}

	//1
	var employee Employee = FullTimeEmployees[0]
	salary, tax := employee.SalaryCalculator(FullTimeEmployees[0].ExtraHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	//2
	employee = FullTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(FullTimeEmployees[1].ExtraHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	//3
	employee = PartTimeEmployee[0]
	salary, tax = employee.SalaryCalculator(PartTimeEmployee[0].PartTimeHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	//4
	employee = PartTimeEmployee[1]
	salary, tax = employee.SalaryCalculator(PartTimeEmployee[0].PartTimeHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	//5
	employee = ShiftEmployee[0]
	salary, tax = employee.SalaryCalculator(ShiftEmployee[0].ShiftHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	//6
	employee = ShiftEmployee[1]
	salary, tax = employee.SalaryCalculator(ShiftEmployee[1].ShiftHours)
	fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
}

type Employee interface {
	SalaryCalculator(hour float64) (salary float64, tax float64)
}
type FullTimeEmployee struct {
	Id           int
	FullName     string
	NationalCode string
	ExtraHours   float64
}
type PartTimeEmployee struct {
	Id            int
	FullName      string
	NationalCode  string
	PartTimeHours float64
}

type ShiftEmployee struct {
	Id           int
	FullName     string
	NationalCode string
	ShiftHours   float64
}

func (employee FullTimeEmployee) SalaryCalculator(extraHour float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHour)*1.4
	tax = TaxRate * salary
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator(Hour float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * Hour
	tax = TaxRate * salary
	salary += tax
	return
}
func (employee ShiftEmployee) SalaryCalculator(Hour float64) (salary float64, tax float64) {
	salary = ShiftSalaryRate * Hour
	tax = TaxRate * salary
	salary += tax
	return
}
