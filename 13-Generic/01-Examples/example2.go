package main

import "fmt"

type Number1 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {
	slc1 := []int{12, 52, 546, 23, 14}
	slc2 := []float64{12.6, 14.46, 52.96, 12.36}
	fmt.Printf("%d\n", Plus(slc1))
	fmt.Printf("%f\n", Plus(slc2))
}
func Plus[T Number1](slc []T) (sum T) {
	for _, i := range slc {
		sum += i
	}
	return
}
