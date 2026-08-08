package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type IOError struct {
	FileName string
	Message  string
	Err      error
}

func main() {
	err := CopyFile1("src.txt", "dst.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Unwrapped error: %s\n", errors.Unwrap(err))
	}
}
func (error *IOError) Unwrap() error {
	return error.Err
}
func (error IOError) Error() string {
	return fmt.Sprintf("IO error occurred :filename:%s  Message:%s  Detail:%s", error.FileName, error.Message, error.Err.Error())
}
func CopyFile1(sourceName, destinationName string) error {
	source, err := os.Open(sourceName)
	if err != nil {
		return &IOError{FileName: sourceName, Message: "during open file could not open source file", Err: err}
	}
	/*defer*/ source.Close()

	destination, err := os.Create(destinationName)
	if err != nil {
		return &IOError{FileName: sourceName, Message: "during open file could not create destination file", Err: err}

	}
	_, err = io.Copy(destination, source)
	if err != nil {
		return &IOError{FileName: sourceName, Message: "during copy file could not copy  file", Err: err}

	}
	//defer destination.Close()
	return nil
}
