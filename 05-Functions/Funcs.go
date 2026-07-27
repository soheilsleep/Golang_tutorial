package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "soheilsleep is a god god god god"
	fmt.Println(strings.Contains(myString, "go12"))
	fmt.Println(strings.ContainsAny(myString, "g56"))
	fmt.Println(strings.Count(myString, "god"))
	fmt.Println(strings.Cut(myString, "god"))

	myStringArray := strings.Split(myString, " ")
	fmt.Println("word count:", len(myStringArray))
	for _, item := range myStringArray {
		fmt.Println(item)
	}
	myStringArray2 := strings.Join(myStringArray, ",")
	println(myStringArray2)
	println(strings.Repeat("iran ", 10))

	println(strings.Replace(myString, "god", "godding", 2))
	println(strings.ReplaceAll(myString, "god", "boy"))

	println(strings.Compare("god", "god"))
	println(strings.Compare("God", "god"))
	println(strings.Compare("god", "GOD"))

	println(strings.EqualFold("god", "god"))

	println(strings.HasPrefix("iran", "ir"))
	println(strings.HasPrefix("iran", "IR"))

	println(strings.HasSuffix("iran", "an"))
	println(strings.HasSuffix("iran", "n"))

	println(strings.Index("iran", "an"))
	println(strings.Index("iran", "n"))

	println(strings.ToLower("irAn"))
	println(strings.ToUpper("iRan"))
	println(strings.Title("iran is a great country"))

	println(strings.Trim("!!iran is a great country!!", "!"))
	println(strings.TrimLeft("!!iran is a great country!!", "!"))
	println(strings.TrimRight("!!iran is a great country!!", "!"))

}
