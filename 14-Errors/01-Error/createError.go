package main

import (
	"errors"
	"fmt"
)

func main() {
	output, err := createErrorMethod1(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}
func createErrorMethod1(number int) (int, error) {
	if number == 0 {
		return 0, errors.New("number is not valid")
	}
	return number * 5, nil
}
func createErrorMethod2(number int) (int, error) {
	if number == 0 {
		return 0, fmt.Errorf("number %d is not valid", number)
	}
	return number * 5, nil
}
