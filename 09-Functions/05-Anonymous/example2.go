package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{12, 63, 74, 58, 2, 3, 45, 64, 21}
	fmt.Printf("numbers: %v\n", numbers)
	sortingFunc := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}
	sort.Slice(numbers, sortingFunc)
	fmt.Printf("numbers: %v\n", numbers)
}
