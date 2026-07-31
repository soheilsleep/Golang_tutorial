package main

import "time"

func main() {
	firstName := "ali"
	names := []string{"soheil", "ali", "erfan", "taha", "taha4", "taha3", "taha2"}

	printFirstNameFunc := func() {
		println(firstName)
	}
	firstName = "soheil"
	printFirstNameFunc()
	//--------------------------
	for i, name := range names {
		go func() {
			name = "*" + name
			println(i, name)
		}()
	}
	time.Sleep(time.Second * 1)

	printFirstNameFunc()
	//--------------------------
	for i, item := range names {
		go func(index int, name string) {
			name = "*" + name
			println(index, name)
		}(i, item)
	}
	time.Sleep(time.Second * 1)
}
