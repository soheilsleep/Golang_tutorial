package ExternalServices

import (
	"fmt"
)

type SmsService struct {
}

//	SendNotify func (s *SmsService) SendMessage(order *entities.Order) {
//		fmt.Printf("Sms sent: %v\n", order)
//	}
func (s *SmsService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}
	fmt.Printf("Sms sent: receiver:%s , Message:%s\n", receiver, message)
	event := logger.Info().Str("notifier", "smsService")
	event.Dict("messageInfo",
		event.CreateDict().
			Str("receiver", receiver).
			Str("message", message),
	).Msg("msg sent")
	return nil
}
func NewSmsService() *SmsService {
	return &SmsService{}
}
