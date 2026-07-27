package main
import (
	"fmt"
	"unsafe"
)

// main is the entry point of the program
// This demonstrates runes and strings in Go, including Unicode support
func main()  {
	// Rune is a 32-bit integer that represents a Unicode code point
	// It's used to handle Unicode characters properly
	char1 := '🤣'	// Unicode emoji
	char2 := '😈'	// Unicode emoji
	char3 := 128525	// Unicode code point number

	// Print rune values and their types
	fmt.Printf("char1 : %d %T \n", char1, char1)
	fmt.Printf("char2 : %d %T \n", char2, char2)
	fmt.Printf("char3 : %c %T \n", char3, char3)

	// String in Go is a sequence of bytes, not characters
	// This means it can't handle Unicode characters properly
	myStr := "Hello soheilsleep 😁💕"
	fmt.Printf("myStr : %s %T, len :%d, size:%d \n", myStr, myStr,len(myStr), unsafe.Sizeof(myStr))
	
	// When iterating over a string, you get bytes, not characters
	// This can cause issues with multi-byte Unicode characters
	for i := 0; i < len(myStr); i++ {
		fmt.Printf("myStr %d: %c %T\n", i, myStr[i], myStr[i])
	}

	// []rune converts a string to a slice of runes
	// This allows proper handling of Unicode characters
	myRune := []rune("Hello soheilsleep 😁💕")
	fmt.Printf("myRune : %v %T, len :%d, size:%d \n", myRune, myRune,len(myRune), unsafe.Sizeof(myRune))

	// When iterating over a rune slice, you get actual characters
	for i := 0; i < len(myRune); i++ {
		fmt.Printf("myStr %d: %c %T\n", i, myRune[i], myRune[i])
	}
}