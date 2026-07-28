package main

import (
	"fmt"
	"strings"
)

func main() {
	var notificationType string // sms, email,push
	println("Please enter a notification type: ")
	fmt.Scanln(&notificationType)

	switch {
	case strings.Contains(notificationType, "sms"):
		println("sms sent")
		fallthrough
	case strings.Contains(notificationType, "email"):
		println("email sent")
		fallthrough

	case strings.Contains(notificationType, "push"):
		println("push sent")
	default:
		println("unknown notification type")
	}
}
