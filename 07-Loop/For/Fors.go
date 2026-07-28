package main

import (
	"fmt"
)

func main() {
	i := 0
	lst1 := []int{12, 13, 14, 54, 65, 65, 12}
	for {
		println(i)
		break
	}
	for i < 10 {
		fmt.Println("hello world", i)
		i++
	}
	for j := 0; j < 10; j++ {
		println("hello world", j)
	}
	for index, item := range lst1 {
		println("Hello world", index, item)
	}
}
