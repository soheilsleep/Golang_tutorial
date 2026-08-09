package ExternalServices

import (
	"fmt"
)

type EmailService struct {
}

//	SendNotify func (e *EmailService) SendMessage(order *entities.Order) *entities.Order {
//		fmt.Printf("Email sent: %v\n", order)
//		return order
//	}
func (e *EmailService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("Email sent: receiver:%s , Message:%s\n", receiver, message)
	event := logger.Info().Str("notifier", "emailService")
	event.Dict("messageInfo",
		event.CreateDict().
			Str("receiver", receiver).
			Str("message", message),
	).Msg("msg sent")
	return nil
}
func NewEmailService() *EmailService {
	return &EmailService{}
}
