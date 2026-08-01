package main

type Employee struct {
	ID            int
	Name          string
	Type          string
	SalaryOfMonth int
}

func main() {
	var employee1 Employee = Employee{1, "soheil", "manager", 15000000}
	SalaryOfYear1 := CalcYearlySalary(employee1)
	SalaryOfYear2 := employee1.CalcYearlySalary()
	println(SalaryOfYear1)
	println(SalaryOfYear2)

}

// CalcYearlySalary Function
func CalcYearlySalary(employee Employee) int {
	return employee.SalaryOfMonth * 12
}

// CalcYearlySalary Method of Employee struct
func (employee *Employee) CalcYearlySalary() int {
	return employee.SalaryOfMonth * 12
}
