package main

import (
	"log"
	"os"
)

var (
	warnLogger  *log.Logger
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	flags := log.Ldate | log.Ltime | log.Lshortfile
	log.SetFlags(flags)
	warnLogger = log.New(file, "WARNING: ", flags)
	errorLogger = log.New(file, "ERROR: ", flags)
	infoLogger = log.New(file, "INFO: ", flags)
	log.SetOutput(file)
}

func main() {
	infoLogger.Println("Hello World")
	sum(1, 6)
}
func sum(a, b int) {
	warnLogger.Println("start of sum")
	println(a + b)
	warnLogger.Println("end of sum")
}
