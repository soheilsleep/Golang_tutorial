package main

import "fmt"

func main() {
	var month int
	var totalDays int = 0

	const monthDays1 int = 30
	const monthDays2 int = 31
	const monthDays3 int = 29

	println("enter the month number")
	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDays += monthDays3
		fallthrough

	case 11:
		totalDays += monthDays1
		fallthrough

	case 10:
		totalDays += monthDays1
		fallthrough

	case 9:
		totalDays += monthDays1
		fallthrough

	case 8:
		totalDays += monthDays1
		fallthrough

	case 7:
		totalDays += monthDays1
		fallthrough

	case 6:
		totalDays += monthDays2
		fallthrough

	case 5:
		totalDays += monthDays2
		fallthrough

	case 4:
		totalDays += monthDays2
		fallthrough

	case 3:
		totalDays += monthDays2
		fallthrough

	case 2:
		totalDays += monthDays2
		fallthrough

	case 1:
		totalDays += monthDays2

	}
	println(totalDays)

}
