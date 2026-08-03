package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	Family    string
	Gender    string
	Height    int
	Weight    int
	HairColor string
}
type PersonOption func(*Person)

func main() {
	person := NewPerson(SetName("James"), SetAge(19), SetFamily("alen"), SetGender("male"), SetHairColor("black"), SetHeight(150), SetWeight(60))
	fmt.Printf("person = %+v\n", person)
}
func NewPerson(options ...PersonOption) *Person {
	person := &Person{Name: "soheil", Age: 20}
	for _, option := range options {
		option(person)
	}
	return person
}
func SetName(name string) PersonOption {
	return func(builder *Person) {
		builder.Name = name
	}
}
func SetAge(age int) PersonOption {
	return func(builder *Person) {
		builder.Age = age
	}
}
func SetGender(gender string) PersonOption {
	return func(builder *Person) {
		builder.Gender = gender
	}
}
func SetFamily(family string) PersonOption {
	return func(builder *Person) {
		builder.Family = family
	}
}
func SetHeight(height int) PersonOption {
	return func(builder *Person) {
		builder.Height = height
	}
}
func SetWeight(weight int) PersonOption {
	return func(builder *Person) {
		builder.Weight = weight
	}
}
func SetHairColor(hairColor string) PersonOption {
	return func(builder *Person) {
		builder.HairColor = hairColor
	}
}
func build() PersonOption {
	return func(builder *Person) {

	}
}
