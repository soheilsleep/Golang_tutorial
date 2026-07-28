package main

import "fmt"

func main() {
	var myArr0 [5]int
	var myArr1 [3]int = [3]int{1, 2}
	myArr2 := [6]int{1, 2}
	myArr3 := [...]int{1, 2}

	myArr0[2] = 20
	myArr1[0] = 20
	myArr2[4] = 20
	myArr3[1] = 20

	fmt.Println(myArr0)
	fmt.Println(myArr1)
	fmt.Println(myArr2)
	fmt.Println(myArr3)

	fmt.Println(len(myArr0))
	fmt.Println(len(myArr1))
	fmt.Println(len(myArr2))
	fmt.Println(len(myArr3))
}
