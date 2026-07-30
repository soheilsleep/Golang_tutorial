package main

func main() {
	myFun := func() {
		println("Hello World")
	}
	myFun()
	println(func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
		}
		return total
	}(1, 2, 3, 4, 5))

	sum := func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
		}
		return total
	}
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
