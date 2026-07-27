package main

import (
	"fmt"
	"unsafe"
	"math/bits"
	
)

// main is the entry point of the program
// This file demonstrates different integer and float types in Go
func main()  {
	// Integer types in Go
	// int: Platform-dependent integer type (usually 32-bit or 64-bit)
	var num int
	// int8: 8-bit signed integer (-128 to 127)
	var num8 int8
	// int16: 16-bit signed integer (-32768 to 32767)
	var num16 int16
	// int32: 32-bit signed integer (-2^31 to 2^31-1)
	var num32 int32
	// int64: 64-bit signed integer (-2^63 to 2^63-1)
	var num64 int64

	// unsafe.Sizeof returns the size in bytes of the variable
	// This is useful for understanding memory usage
	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num))
	fmt.Printf("num8  %d bytes \n" , unsafe.Sizeof(num8))
	fmt.Printf("num16  %d bytes \n" , unsafe.Sizeof(num16))
	fmt.Printf("num32  %d bytes \n" , unsafe.Sizeof(num32))
	fmt.Printf("num64  %d bytes \n" , unsafe.Sizeof(num64))

	// bits.UintSize returns the size of an int in bits (usually 32 or 64)
	var a = bits.UintSize
	fmt.Println(a)

	// Float types in Go
	// float32: 32-bit floating-point number
	var fnum float32 = 10.3
	// float64: 64-bit floating-point number (default for decimal numbers)
	var fnum8 float64 

	fmt.Printf("fnum  %d bytes \n" , unsafe.Sizeof(fnum))
	fmt.Printf("fnum8  %d bytes \n" , unsafe.Sizeof(fnum8))

	
}