package main

import (
	"fmt"

	jalaali "github.com/jalaali/go-jalaali"
	"github.com/soheilsleep/ModuleExample/services"
)

func main() {
	fmt.Println("Hello World")
	year, month, day, error := jalaali.ToGregorian(1399, 13, 12)
	if error == nil {
		fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
	} else {
		fmt.Printf("Error: %v\n", error)
	}
	var service services.TestService = services.TestService{}
	fmt.Printf("%v\n", service)
}
