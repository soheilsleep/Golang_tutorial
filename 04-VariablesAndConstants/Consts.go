package main

import "fmt"

func main(){
	const (
		name = "soheilsleep"
		number = 26
		city = "ahvaz"
		pi = 3.14
		)

		fmt.Printf("name : %s number : %d city:%s pi: %f \n", name, number, city, pi)

		const GoogleBaseUrl = "https://www.google.com"
		const MapUrl = "/maps"

		fmt.Println(GoogleBaseUrl, MapUrl)
		
}