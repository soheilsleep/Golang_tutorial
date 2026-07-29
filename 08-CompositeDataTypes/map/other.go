package main

import "fmt"

type person struct {
	name   string
	family string
	age    int
}

func main() {
	//insert
	persons := make(map[string]person)
	personSlice := []string{}
	persons["1234567890"] = person{name: "soheil", family: "sleep", age: 26}
	personSlice = append(personSlice, "1234567890")
	persons["8948519842"] = person{name: "ahmad", family: "kazemi", age: 50}
	personSlice = append(personSlice, "8948519842")
	persons["7816842692"] = person{name: "taha", family: "dehnavi", age: 7}
	personSlice = append(personSlice, "7816842692")
	persons["8746518684"] = person{name: "saeed", family: "balooch", age: 7}
	personSlice = append(personSlice, "8746518684")
	persons["5642198458"] = person{name: "mehran", family: "gol payegan", age: 7}
	personSlice = append(personSlice, "5642198458")
	persons["5741894965"] = person{name: "ghasem", family: "rezaee", age: 7}
	personSlice = append(personSlice, "5741894965")

	fmt.Printf("%d\n", len(persons))
	for _, person := range personSlice {
		fmt.Printf("%v\n", persons[person])
	}
	//interview questions:
	//1. References type or Value type
	//2. Copy map
	persons2 := make(map[string]person)
	for key, value := range persons {
		persons2[key] = value
	}

}
