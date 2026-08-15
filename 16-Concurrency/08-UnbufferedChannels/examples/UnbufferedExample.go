package examples

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func UnbufferedExample() {
	now := time.Now()
	baseUrl := "https://jsonplaceholder.typicode.com/todos/"
	content := make(chan string)
	for i := range 100 {
		go GetHttpRequestUnbuffered(content, baseUrl, i+1)
	}
	for range 100 {
		response := <-content
		fmt.Println(response)
	}
	println("Time:", time.Since(now).Milliseconds())

}
func GetHttpRequestUnbuffered(content chan<- string, baseUrl string, index int) {

	Response, err := http.Get(baseUrl + strconv.Itoa(index))
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(Response.Body)
	responseBody, err := io.ReadAll(Response.Body)
	if err != nil {
		panic(err)
	}
	PrintlnWithTime("Before set content")
	content <- string(responseBody)
	PrintlnWithTime("After set content")
}
func PrintlnWithTime(args ...any) {
	fmt.Printf("time:%s, %v\n", time.Now().Format(time.RFC3339Nano), args)
}
