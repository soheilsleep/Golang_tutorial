package examples

import (
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RequestResponse struct {
	Url          string
	Index        int
	ResponseBody string
}

func BufferedExample() {
	total := 100
	wg := new(sync.WaitGroup)
	now := time.Now()
	content := make(chan RequestResponse, total)
	wg.Add(total)
	for i := range total {
		content <- RequestResponse{
			Url:   "https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(i+1),
			Index: i,
		}

		go GetHttpRequestBuffered(content, wg)
	}
	wg.Wait()
	close(content)
	for item := range content {
		PrintlnWithTime(item)
	}
	PrintlnWithTime("Time:", time.Since(now).Milliseconds())

}
func GetHttpRequestBuffered(content chan RequestResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	responseRequest := <-content
	Response, err := http.Get(responseRequest.Url)
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
	content <- RequestResponse{Url: responseRequest.Url, Index: responseRequest.Index, ResponseBody: string(responseBody)}
	PrintlnWithTime("After set content")
}
