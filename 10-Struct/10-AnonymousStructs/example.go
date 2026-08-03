package main

import "fmt"

type Person struct {
	int
	string
	float64
}

func main() {
	// Api call and get response
	apiResponse := struct {
		ResultCode        int
		ResultMsg         string
		TransactionAmount float64
		TransactionTime   string
	}{
		ResultCode:        0,
		ResultMsg:         "success",
		TransactionAmount: 1.0,
		TransactionTime:   "2020-09-07T00:00:00+09:00",
	}
	fmt.Printf("%+v\n", apiResponse)
	person := Person{1, "soheil", 12.66}
	fmt.Printf("%+v\n", person)
}
