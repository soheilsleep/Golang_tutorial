package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	function1(nums, 7)
	divide(10, 5)
	divide(10, 0)
}

func function1(numbers []int, index int) {
	//1- index out of range
	//fmt.Println(numbers[index])
	//2-Panic
	//panic("Something bad happened"
	//3-Log
	//log.Fatal("fatal")
}
func divide(a, b int) {
	defer RecoverFunc()
	fmt.Printf("start of divide:%d %d \n", a, b)
	fmt.Printf("result:%d \n", a/b)
	fmt.Printf("end of divide:%d %d \n", a, b)
}
func RecoverFunc() {
	if err := recover(); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Recovered from panic")
		//debug.PrintStack()
	}
}
