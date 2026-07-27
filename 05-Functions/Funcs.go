package main

import (
	"fmt"
	"strings"
)

// main is the entry point of the program
// This demonstrates various string manipulation functions in Go
func main() {
	// String manipulation functions are part of the strings package
	// These functions are safe to use and don't modify the original string
	
	myString := "soheilsleep is a god god god god"
	
	// Check if a substring exists in the string
	// Returns true if the substring is found, false otherwise
	fmt.Println(strings.Contains(myString, "go12"))
	
	// Check if any of the specified characters exist in the string
	fmt.Println(strings.ContainsAny(myString, "g56"))
	
	// Count the number of non-overlapping occurrences of a substring
	fmt.Println(strings.Count(myString, "god"))
	
	// Cut splits the string at the first occurrence of the separator
	// Returns the two parts and a boolean indicating if the separator was found
	fmt.Println(strings.Cut(myString, "god"))

	// Split a string into a slice of substrings based on a separator
	myStringArray := strings.Split(myString, " ")
	fmt.Println("word count:", len(myStringArray))
	for _, item := range myStringArray {
		fmt.Println(item)
	}
	
	// Join a slice of strings into a single string with a separator
	myStringArray2 := strings.Join(myStringArray, ",")
	println(myStringArray2)
	
	// Repeat a string a specified number of times
	println(strings.Repeat("iran ", 10))

	// Replace occurrences of a substring with another substring
	// The third parameter specifies the number of replacements
	println(strings.Replace(myString, "god", "godding", 2))
	
	// Replace all occurrences of a substring with another substring
	println(strings.ReplaceAll(myString, "god", "boy"))

	// Compare two strings lexicographically
	// Returns: -1 if s1 < s2, 0 if s1 == s2, 1 if s1 > s2
	println(strings.Compare("god", "god"))
	println(strings.Compare("God", "god"))
	println(strings.Compare("god", "GOD"))

	// Compare two strings case-insensitively
	println(strings.EqualFold("god", "god"))

	// Check if a string starts with a prefix
	println(strings.HasPrefix("iran", "ir"))
	println(strings.HasPrefix("iran", "IR"))

	// Check if a string ends with a suffix
	println(strings.HasSuffix("iran", "an"))
	println(strings.HasSuffix("iran", "n"))

	// Find the first index of a substring
	// Returns the index of the first occurrence, or -1 if not found
	println(strings.Index("iran", "an"))
	println(strings.Index("iran", "n"))

	// Convert a string to lowercase
	println(strings.ToLower("irAn"))
	
	// Convert a string to uppercase
	println(strings.ToUpper("iRan"))
	
	// Capitalize the first letter of each word
	println(strings.Title("iran is a great country"))

	// Trim characters from both ends of a string
	println(strings.Trim("!!iran is a great country!!", "!"))
	
	// Trim characters from the left side of a string
	println(strings.TrimLeft("!!iran is a great country!!", "!"))
	
	// Trim characters from the right side of a string
	println(strings.TrimRight("!!iran is a great country!!", "!"))

}