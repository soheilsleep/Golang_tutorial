package main

import (
	"sync"
	"sync/atomic"
)

type Employee struct {
	ID     int
	Salary int64
}

func main() {
	wg := sync.WaitGroup{}
	//mx := sync.Mutex{}
	var TotalBalance int64 = 25_500_000_000
	employeeSalaryList := []Employee{}
	wg.Add(5000)
	for i := 0; i < 5_000; i++ {
		employeeSalaryList = append(employeeSalaryList, Employee{ID: i, Salary: GetRandomNumber()})
	}
	for _, employee := range employeeSalaryList {
		go func(employee Employee) {
			defer wg.Done()
			if employee.Salary < TotalBalance {
				// Method 1
				//mx.Lock()
				//TotalBalance -= employee.Salary
				//mx.Unlock()

				//Method 2
				atomic.AddInt64(&TotalBalance, -employee.Salary)
			}

		}(employee)

	}
	wg.Wait()
	println(TotalBalance)
}
func GetRandomNumber() int64 {
	return 5_000_000
}
