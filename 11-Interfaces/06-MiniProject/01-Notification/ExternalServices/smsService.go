package ExternalServices

import (
	"fmt"
	"notification/entities"
)

type SmsService struct {
}

func (s *SmsService) SendMessage(order *entities.Order) {
	fmt.Printf("Sms sent: %v\n", order)
}
func (s *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("Sms sent: receiver:%s , Message:%s\n", receiver, message)
}
func NewSmsService() *SmsService {
	return &SmsService{}
}
