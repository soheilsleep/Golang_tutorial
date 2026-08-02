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
		{Employee: Employee{Id: 1, NationalCode: "1234567890", FullName: "soheilsleep"}, ExtraHours: 40},
		{Employee: Employee{Id: 2, FullName: "taha dehnavi", NationalCode: "1548628460"}, ExtraHours: 45},
	}
	PartTimeEmployee := []PartTimeEmployee{
		{Employee: Employee{Id: 3, FullName: "reza mohamadi", NationalCode: "1985181580"}, PartTimeHours: 100},
		{Employee: Employee{Id: 4, FullName: "ali gorbay", NationalCode: "4518456150"}, PartTimeHours: 140},
	}
	ShiftEmployee := []ShiftEmployee{
		{Employee: Employee{Id: 5, FullName: "hediye mohamadi", NationalCode: "7456274596"}, ShiftHours: 50},
		{Employee: Employee{Id: 6, FullName: "asal gorbay", NationalCode: "1462984573"}, ShiftHours: 60},
	}
	var employees []EmployeeSalaryCalculator = []EmployeeSalaryCalculator{}
	for _, employee := range FullTimeEmployees {
		employees = append(employees, employee)
	}
	for _, employee := range PartTimeEmployee {
		employees = append(employees, employee)
	}
	for _, employee := range ShiftEmployee {
		employees = append(employees, employee)
	}
	for _, employee := range employees {
		salary, tax := employee.SalaryCalculate()
		fmt.Printf("Employee(%T): %v\nSalary: %f\n Tax: %f\n", employee, employee, salary, tax)
	}
}

type EmployeeSalaryCalculator interface {
	SalaryCalculate() (salary float64, tax float64)
}
type Employee struct {
	Id           int
	FullName     string
	NationalCode string
}
type FullTimeEmployee struct {
	Employee
	ExtraHours float64
}
type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

type ShiftEmployee struct {
	Employee
	ShiftHours float64
}

func (employee FullTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = TaxRate * salary
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = TaxRate * salary
	salary += tax
	return
}
func (employee ShiftEmployee) SalaryCalculate() (salary float64, tax float64) {
	salary = ShiftSalaryRate * employee.ShiftHours
	tax = TaxRate * salary
	salary += tax
	return
}
