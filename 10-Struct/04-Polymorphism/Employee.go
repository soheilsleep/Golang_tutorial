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
	var employees []Employee = []Employee{}
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
		salary, tax := employee.SalaryCalculator()
		fmt.Printf("Employee(%T): %v\nSalary: %f\n Tax: %f\n", employee, employee, salary, tax)
	}
}

type Employee interface {
	SalaryCalculator() (salary float64, tax float64)
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

func (employee FullTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = TaxRate * salary
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = TaxRate * salary
	salary += tax
	return
}
func (employee ShiftEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = ShiftSalaryRate * employee.ShiftHours
	tax = TaxRate * salary
	salary += tax
	return
}
