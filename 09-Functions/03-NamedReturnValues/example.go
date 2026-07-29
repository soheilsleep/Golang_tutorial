package main

import "fmt"

func main() {
	price, finalprice, tax := CalculateRoomPrice("double", 3, 2)
	fmt.Printf("price is %d and final price is :%d  and tax is %.2f", price, finalprice, tax)
}
func CalculateRoomPrice(roomType string, nights int, personCount int) (price int, finalPrice int, tax float64) {
	switch roomType {
	case "twin":
		price = nights * 120000 * personCount
	case "double":
		price = nights * 200000 * personCount
	case "suite":
		price = nights * 500000 * personCount

	}
	tax = float64(price) * 0.09
	finalPrice = price + int(tax)
	return
}
