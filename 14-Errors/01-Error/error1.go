package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://mykettt.ir/")
	if err != nil {
		fmt.Println("an error occurred on get request")
		return
	}
	println(response.Status)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("an error occurred on reading response body")
		return
	}
	fmt.Printf("%s\n", responseBody)
}
