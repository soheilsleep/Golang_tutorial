package main

import (
	"fmt"
	"time"
)

func main() {
	numChannel1 := make(chan int)
	go SendDataToChannel(numChannel1)
	receivedData := <-numChannel1
	PrintlnWithTime("Received num:", receivedData)
	receivedData = <-numChannel1
	PrintlnWithTime("Received num:", receivedData)
	time.Sleep(2 * time.Second)
}

func SendDataToChannel(numChannel1 chan int) {
	PrintlnWithTime("Before sending 1 to channel")
	numChannel1 <- 1
	PrintlnWithTime("After sending 1 to channel")

	PrintlnWithTime("Before sending 2 to channel")
	numChannel1 <- 2
	PrintlnWithTime("After sending 2 to channel")

	PrintlnWithTime("Before sending 3 to channel")
	numChannel1 <- 3
	PrintlnWithTime("After sending 3 to channel")

}
func ReceiveDataFromChannel(numChannel1 chan int) {}

func PrintlnWithTime(args ...any) {
	fmt.Printf("Time:%s ,%v\n", time.Now().Format(time.RFC3339Nano), args)
}
