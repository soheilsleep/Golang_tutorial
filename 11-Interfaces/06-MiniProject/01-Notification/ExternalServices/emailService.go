package ExternalServices

import (
	"fmt"
	"notification/entities"
)

type EmailService struct {
}

func (e *EmailService) SendMessage(order *entities.Order) *entities.Order {
	fmt.Printf("Email sent: %v\n", order)
	return order
}
func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email sent: receiver:%s , Message:%s\n", receiver, message)
}
func NewEmailService() *EmailService {
	return &EmailService{}
}
