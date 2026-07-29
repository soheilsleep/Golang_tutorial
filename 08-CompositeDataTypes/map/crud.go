package main

import "fmt"

type Person struct {
	name   string
	family string
	age    int
}

func main() {
	//insert
	persons := make(map[string]Person)
	persons["1234567890"] = Person{name: "soheil", family: "sleep", age: 26}
	persons["8948519842"] = Person{name: "ahmad", family: "kazemi", age: 50}
	persons["7816842692"] = Person{name: "taha", family: "dehnavi", age: 7}

	fmt.Printf("%v\n", persons)
	//edit
	persons["1234567890"] = Person{name: "ali", family: "yaghubi", age: 26}
	fmt.Printf("%v\n", persons)
	//delete
	delete(persons, "8948519842")
	fmt.Printf("%v\n", persons)

	//read
	Ali, ok := persons["1244567890"]
	if ok {
		fmt.Printf("%v\n", Ali)
	} else {
		fmt.Printf("not exist\n")
	}

}
