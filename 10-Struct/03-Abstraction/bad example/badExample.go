package main

import (
	"fmt"
)

const (
	BaseSalary       = 6600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate  = 80000
	TaxRate          = 0.09
)

func main1() {
	employees := []Employee{
		{1, "soheilsleep", "1234567890", "FullTime", 40},
		{2, "ali gorbay", "4518456150", "PartTime", 120},
		{3, "taha dehnavi", "1548628460", "FullTime", 45},
		{3, "reza mohamadi", "1985181580", "Shift", 160},
	}
	for _, employee := range employees {
		salary, tax := employee.SalaryCalculator(float64(employee.Hours))
		fmt.Printf("Employees: %v\n  Salary : %f\n tax: %f\n", employee, salary, tax)
	}
}

type Employee struct {
	Id           int
	FullName     string
	NationalCode string
	Type         string
	Hours        float64
}

func (employee Employee) FullTimeSalaryCalculator(extraHour float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHour)*1.4
	tax = TaxRate * salary
	salary += tax
	return
}

func (employee Employee) PartTimeSalaryCalculator(Hour float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * Hour
	tax = TaxRate * salary
	salary += tax
	return
}
func (employee Employee) ShiftSalaryCalculator(Hour float64) (salary float64, tax float64) {
	salary = ShiftSalaryRate * Hour
	tax = TaxRate * salary
	salary += tax
	return
}
func (employee Employee) SalaryCalculator(Hour float64) (salary float64, tax float64) {
	if employee.Type == "FullTime" {
		return employee.FullTimeSalaryCalculator(Hour)
	} else if employee.Type == "PartTime" {
		return employee.PartTimeSalaryCalculator(Hour)
	} else {
		return employee.ShiftSalaryCalculator(Hour)
	}
	return
}
