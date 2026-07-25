package main
import (
	"fmt"
	"unsafe"
	
)
func main()  {
	var num int
	var num8 int8
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num))
	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num8))
	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num16))
	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num32))
	fmt.Printf("num  %d bytes \n" , unsafe.Sizeof(num64))

	var a = bits.UintSize
	fmt.Println(a)
}