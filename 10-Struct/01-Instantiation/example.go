package main

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
	Age       uint
}
type PersonOption struct {
	FirstName string
	Lastname  string
	Age       uint
}

func main() {
	//1
	Person1 := Person{"reza", "konian", 3}
	fmt.Println(Person1)
	//2
	Person2 := new(Person)
	Person2.FirstName = "soheil"
	Person2.Lastname = "sleep"
	Person2.Age = 26
	fmt.Println(Person2)
	//3
	Person3 := &Person{FirstName: "Ali", Lastname: "yaghubi", Age: 266}
	fmt.Println(Person3)
	//4
	Person4 := NewPerson("taha", "deh", 12)
	fmt.Println(Person4)
	//5
	Person5 := NewPersonWithOption(PersonOption{FirstName: "milad", Lastname: "miri", Age: 12})
	fmt.Println(Person5)
}
func NewPerson(FirstName string, Lastname string, Age uint) *Person {
	if len(Lastname) < 4 {
		return nil
	}
	return &Person{FirstName: FirstName, Lastname: Lastname, Age: Age}
}
func NewPersonWithOption(option PersonOption) *Person {
	if len(option.Lastname) < 4 {
		return nil
	}
	return &Person{FirstName: option.FirstName, Lastname: option.Lastname, Age: option.Age}
}
