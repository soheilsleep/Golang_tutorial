package main

import "fmt"

func main() {
	myArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	//mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//mySlice1 := make([]int, 8)
	//mySlice2 := make([]int, 8, 16)
	slc := myArray[2:6]
	myArray[0] = 100
	//fmt.Printf("%v\n", slc)
	//fmt.Printf("%v\n", myArray)
	//
	//fmt.Println("slc length:", len(slc))
	//fmt.Println("myArray length:", len(myArray))
	//fmt.Println("slc cap:", cap(slc))
	//fmt.Println("myArray cap:", cap(myArray))

	slc = append(slc, 15)
	slc = append(slc, 17)
	slc = append(slc, 19)
	myArray[2] = 200

	fmt.Printf("%v\n", myArray)
	fmt.Printf("%v\n", slc)

	fmt.Println("slc length:", len(slc))
	fmt.Println("myArray length:", len(myArray))
	fmt.Println("slc cap:", cap(slc))
	fmt.Println("myArray cap:", cap(myArray))
}
