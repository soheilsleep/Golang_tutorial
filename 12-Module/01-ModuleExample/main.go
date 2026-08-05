package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"
)

func main() {
	fmt.Println("Hello World")
	year, month, day, error := jalaali.ToGregorian(1399, 1, 1)
	if error == nil {
		fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
	} else {
		fmt.Printf("Error: %v\n", error)
	}
}
