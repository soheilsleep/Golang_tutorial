package examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SelectExample1() {
	resource1 := make(chan string)
	resource2 := make(chan string)

	go GetResponse(
		resource1,
		"https://sportscore.com/api/widget/matches/?sport=football&limit=50",
	)

	go GetResponse(
		resource2,
		"https://soccerresults.se/api/data.php?action=today",
	)

	select {
	case result := <-resource1:
		fmt.Println("Result from SportScore:")
		fmt.Println(result)

	case result := <-resource2:
		fmt.Println("Result from SoccerResults:")
		fmt.Println(result)
	case <-time.After(time.Second * 10):
		fmt.Println("Timed out")
	default:
		fmt.Println("No results")
	}

	PrintlnWithTime("End")
}

func GetResponse(content chan<- string, url string) {
	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	destination := &bytes.Buffer{}

	if err := json.Indent(destination, responseBody, "", "  "); err != nil {
		panic(err)
	}

	content <- destination.String()
}
func SelectExample2() {
	resource1 := make(chan string)
	resource2 := make(chan string)
	var result string
	go GetResponse(
		resource1,
		"https://sportscore.com/api/widget/matches/?sport=football&limit=50",
	)

	go GetResponse(
		resource2,
		"https://soccerresults.se/api/data.php?action=today",
	)
	for {
		select {
		case result = <-resource1:
			fmt.Println("Result from SportScore:")
			fmt.Println(result)
			return

		case result = <-resource2:
			fmt.Println("Result from SoccerResults:")
			fmt.Println(result)
			return
		case <-time.After(time.Second * 10):
			fmt.Println("Timed out")
			return
		default:
			fmt.Println("No results")
		}
	}
}

func PrintlnWithTime(args ...any) {
	fmt.Printf(
		"Time:%s, %v\n",
		time.Now().Format(time.RFC3339Nano),
		args,
	)
}
