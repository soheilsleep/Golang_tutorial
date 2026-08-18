package examples

import (
	"encoding/json"
)

func ExampleUnmarshal1() {
	person1json := []byte(`{"first_name":"John Doe","last_name":"Smith","age":45}`)
	var person1 = Person{}
	err := json.Unmarshal(person1json, &person1)
	if err != nil {
		panic(err)
	}
	println(person1.Name)
	println(person1.Family)
	println(person1.Age)
}
