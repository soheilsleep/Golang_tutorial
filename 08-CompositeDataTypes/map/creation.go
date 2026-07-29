package main

import "fmt"

func main() {
	//map creation
	names := make(map[string]string)
	names1 := map[string]string{}
	var names2 map[string]string = map[string]string{} //bullshit
	names["1234567890"] = "soheil sleep"
	names1["1234567890"] = "soheil sleep"
	names2["1234567890"] = "soheil sleep"

	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", names1)
	fmt.Printf("%v\n", names2)
}
