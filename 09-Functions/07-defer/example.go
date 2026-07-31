package main

import (
	"io"
	"os"
)

func main() {
	CopyFile("C:\\soheilsleep\\progeraming\\Golang\\projects\\Golang_tutorial\\09-Functions\\07-defer\\destination.txt", "C:\\soheilsleep\\progeraming\\Golang\\projects\\Golang_tutorial\\09-Functions\\07-defer\\source.txt")
}

func CopyFile(destinationName, sourceName string) (written int64, err error) {
	source, err := os.Open(sourceName)
	if err != nil {
		return
	}
	defer source.Close()
	destination, err := os.Create(destinationName)
	if err != nil {
		return
	}
	defer destination.Close()
	return io.Copy(destination, source)
}
