package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {
	x := Sum(12, 52)
	y := Sum(12.6, 56.8265520)
	fmt.Println(x)
	fmt.Println(y)
}

func SumInt(a, b int) int {
	return a + b
}
func SumFloat(a, b float64) float64 {
	return a + b
}
func Sum[T Number](a, b T) T {
	return a + b
}
