package main
import "encoding/json"

// Order represents a simple order structure
// It contains an ID, price, and status
type Order struct {
    Id int
    Price int
    Status OrderStatus
}

// OrderStatus is a custom type based on int
// This is a common pattern in Go for creating enums
type OrderStatus int

// Constants defining different order statuses
// These are similar to enum values in other languages
const (
	Created = 0
	Processing = 1
	WaitForPayment = 2
	PaymentDone = 3
	Issue = 4
	Refund = 5
	
)

// main is the entry point of the program
// This demonstrates how to use custom types and constants in Go
func main()  {
	// Create three order instances with different statuses
	order1 := Order{Id:1 , Price:100 , Status: Created}
	order2 := Order{Id:2 , Price:200 , Status: PaymentDone}
	order3 := Order{Id:3 , Price:300 , Status: Issue}

	// Marshal orders to JSON format
	// json.Marshal converts Go structs to JSON
	order1Json, _ := json.Marshal(order1)
	order2Json, _ := json.Marshal(order2)
	order3Json, _ := json.Marshal(order3)
	
	// Print the JSON representations
   println(string (order1Json))
   println(string (order2Json))
   println(string (order3Json))


	// These lines would print the numeric values of the constants
	// Uncomment to see the values: 0, 1, 2, 3, 4, 5
	// println(Created)
	// println(Processing)
	// println(WaitForPayment)
	// println(PaymentDone)
	// println(Issue)
	// println(Refund)
}