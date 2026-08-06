package main

import (
	"GenericList/Generics"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	GenericInt()
	GenericString()
	GenericPerson()
}

func GenericInt() {
	list1 := Generics.List[int]{Items: []int{}}
	list1.Add(10)
	list1.Add(20)
	list1.Add(30)
	fmt.Printf("list1: %v\n", list1)
	list1.InsertAt(2, 40)
	fmt.Printf("list1: %v\n", list1)
	list1.RemoveAt(1)
	fmt.Printf("list1: %v\n", list1)
	list1.Remove(30)
	fmt.Printf("list1: %v\n", list1)
}
func GenericString() {
	list2 := Generics.List[string]{Items: []string{}}
	list2.Add("soheil")
	list2.Add("ali")
	list2.Add("reza")
	list2.Add("mohsen")
	fmt.Printf("list2: %v\n", list2)
	list2.RemoveAt(2)
	fmt.Printf("list2: %v\n", list2)
	list2.InsertAt(3, "HajAli")
	fmt.Printf("list2: %v\n", list2)
}

func GenericPerson() {
	list3 := Generics.List[Person]{Items: []Person{}}
	list3.Add(Person{"John", 20})
	list3.Add(Person{"Jane", 30})
	list3.Add(Person{"Eli", 22})
	fmt.Printf("list3: %v\n", list3)
	list3.Remove(Person{"John", 20})
	fmt.Printf("list3: %v\n", list3)
	list3.InsertAt(2, Person{"Ahamd", 12})
	fmt.Printf("list3: %v\n", list3)
}
