package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	changeNumber(numbers)
	fmt.Printf("%v\n", numbers)
	addItem(&numbers)
	fmt.Printf("%v\n", numbers)
}
func changeNumber(numbers []int) {
	for i, _ := range numbers {
		numbers[i] = numbers[i] * 3
	}

}
func addItem(numbers *[]int) {
	*numbers = append(*numbers, 6)
}
