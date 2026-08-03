package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"first_name"`
	Age    int    `json:"age,omitempty"`
	Family string `json:"last_name"`
	Gender string `json:"gender"`
	Height int    `json:"-"`
	Weight int    `json:"weight"`
}

func main() {
	person := Person{"John", 42, "Jim", "Male", 125, 25}
	personJson, _ := json.MarshalIndent(person, "", "	")
	fmt.Println(string(personJson))

}
