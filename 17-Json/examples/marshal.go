package examples

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"first_name"`
	Family string `json:"last_name"`
	Age    int    `json:"age,omitempty"`
}

func ExampleMarshal1() {
	person1 := Person{
		Name:   "John Doe",
		Family: "Smith",
		Age:    42,
	}
	person2 := Person{
		Name:   "Elizabeth",
		Family: "Allen",
		Age:    19,
	}
	person1Json, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}
	person2Json, err := json.Marshal(person2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(person1Json))
	fmt.Println(string(person2Json))
}
func ExampleMarshal2() {
	person1 := Person{
		Name:   "John Doe",
		Family: "Smith",
		Age:    42,
	}
	person2 := Person{
		Name:   "Elizabeth",
		Family: "Allen",
		Age:    19,
	}
	person3 := Person{
		Name:   "soheil",
		Family: "sleep",
		Age:    0,
	}
	persons := []Person{person1, person2, person3}
	personsJson, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(personsJson))
}
