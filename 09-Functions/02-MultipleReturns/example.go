package main

import "fmt"

func main() {
	order1, tax1 := CalculateRoomPrice("double", 3, 2)
	fmt.Printf("price is %d and tax is %.2f", order1, tax1)
}
func CalculateRoomPrice(roomType string, nights int, personCount int) (int, float64) {
	var price int
	var tax float64
	switch roomType {
	case "twin":
		price = nights * 120000 * personCount
	case "double":
		price = nights * 200000 * personCount
	case "suite":
		price = nights * 500000 * personCount

	}
	tax = float64(price) * 0.09
	return price, tax
}
