package main

import "fmt"

func main() {
	printLog(12, "soheil", 45, 45.6, true)
}

func printLog(logs ...interface{}) {
	for i, log := range logs {
		fmt.Printf("%d: %v\n", i, log)
	}
	return
}
