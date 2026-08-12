package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList = []string{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		//wg.Add(1)
		go GetTodo(i+1, &wg)
	}
	wg.Wait()
	fmt.Printf("%v\n", TodoList)
}
func GetTodo(Id int, wg *sync.WaitGroup) {
	//https://jsonplaceholder.typicode.com/todos
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(Id), wg)
}
func GetUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	responseBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	TodoList = append(TodoList, string(responseBody))

}
