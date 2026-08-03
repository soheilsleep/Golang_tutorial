package main

import "fmt"

type Creature interface {
	Eat()
	Sleep()
	Walk()
}

type Human interface {
	Creature
	Think()
	Speak()
}
type Animal interface {
	Creature
}
type Employee struct {
	name string
	age  int
}
type Dog struct {
	name string
}

func main() {
	employee := &Employee{"soheil", 12}
	dog := &Dog{"erfan"}
	var animal Animal = dog
	var human Human = employee
	human.Speak()
	human.Eat()
	human.Sleep()
	human.Walk()
	human.Think()

	animal.Walk()
	animal.Eat()
	animal.Sleep()

}
func (dog Dog) Eat() {
	fmt.Println("dog is eating")
}
func (dog Dog) Sleep() {
	fmt.Println("dog is sleeping")
}
func (dog Dog) Walk() {
	fmt.Println("dog is walking")
}
func (employee Employee) Eat() {
	fmt.Println("employee is eating")
}
func (employee Employee) Sleep() {
	fmt.Println("employee is sleeping")
}
func (employee Employee) Walk() {
	fmt.Println("employee is walking")
}
func (employee Employee) Think() {
	fmt.Println("employee is thinking")
}
func (employee Employee) Speak() {
	fmt.Println("employee is speaking")
}
