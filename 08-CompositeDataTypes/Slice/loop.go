package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"john", "paul", "george", "ringo"}

	for _, item := range names {
		item = strings.ToUpper(item)
	}
	fmt.Println(names)
	//for slice we should use index to change them in a for not the item
	for i, item := range names {
		names[i] = strings.ToUpper(item)
	}
	fmt.Println(names)
}
