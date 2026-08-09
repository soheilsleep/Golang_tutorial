package main

import (
	"time"
)

func main() {
	value := 0
	go function1()
	go function2()
	go function3()
	go func() {
		// lock
		value++ //1
		// unlock
	}()
	go func() {
		// lock
		value += 2 //3
		//unlock
	}()
	go func() {
		//lock
		value += 3 //6
		//unlock
	}()
	println(value)
	time.Sleep(time.Second)
}

func function1() {
	println("Hello World1")
}
func function2() {
	println("Hello World2")
}
func function3() {
	println("Hello World3")
}

// race condition
