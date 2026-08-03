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
type PersonBuilder struct {
	Person
}

func main() {
	//1
	person := Person{"soheil", 12, "kazemi", "boy", 125, 100, "blonde"}
	fmt.Printf("person = %+v\n", person)
	//2 Builder pattern
	personBuilder := &PersonBuilder{}
	person2 := personBuilder.SetName("soheil").SetAge(12).SetWeight(25).SetHeight(150).SetHairColor("blue").SetGender("male").SetFamily("kazemi").Build()
	fmt.Printf("person2 = %+v\n", person2)
}
func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}
func (builder *PersonBuilder) SetAge(age int) *PersonBuilder {
	builder.Age = age
	return builder
}
func (builder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	builder.Gender = gender
	return builder
}
func (builder *PersonBuilder) SetFamily(family string) *PersonBuilder {
	builder.Family = family
	return builder
}
func (builder *PersonBuilder) SetHeight(height int) *PersonBuilder {
	builder.Height = height
	return builder
}
func (builder *PersonBuilder) SetWeight(weight int) *PersonBuilder {
	builder.Weight = weight
	return builder
}
func (builder *PersonBuilder) SetHairColor(hairColor string) *PersonBuilder {
	builder.HairColor = hairColor
	return builder
}
func (builder *PersonBuilder) Build() *Person {
	return &builder.Person
}
