package main

import "fmt"

func main() {
	sum, multiply := Calculator(10, 20, 36, 77)
	fmt.Println("Sum:", sum, "multiple:", multiply)

}

func Calculator(numbers ...int) (sum int, multiplier int) {
	multiplier = 1
	for _, number := range numbers {
		sum += number
		multiplier *= number
	}
	return
}
