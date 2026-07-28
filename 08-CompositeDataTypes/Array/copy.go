package main

import "fmt"

func main() {
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	numbers2 := &numbers
	numbers[0] = 100
	println(&numbers)
	println(&numbers2)

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numbers2)
	changeValue(&numbers)
	changeValue(numbers2)
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numbers2)

}
func changeValue(myArr *[8]int) {
	myArr[0] = 333
	myArr[1] = 999
	myArr[2] = 666

}
