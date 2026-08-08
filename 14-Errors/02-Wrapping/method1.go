package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	err := CopyFile("src.txt", "dst.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func CopyFile(sourceName, destinationName string) error {
	source, err := os.Open(sourceName)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %w", err)
	}
	/*defer*/ source.Close()

	destination, err := os.Create(destinationName)
	if err != nil {
		return fmt.Errorf("during copy file couldn't open destination file %w", err)

	}
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("writing to destination file failed: %w", err)
	}
	//defer destination.Close()
	return nil
}
