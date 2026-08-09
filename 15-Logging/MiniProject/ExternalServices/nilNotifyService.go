package ExternalServices

import (
	"fmt"
	"notification/core"
)

var logger = core.NewFileLogger()

type NilNotifyService struct {
}

func (n NilNotifyService) SendNotify(receiver string, message string) error {
	if receiver == "" {
		return fmt.Errorf("receiver is empty")
	}

	fmt.Printf("NilNotifyService(%s,%s)\n", receiver, message)
	event := logger.Info().Str("notifier", "NilNotifyService")
	event.Dict("messageInfo",
		event.CreateDict().
			Str("receiver", receiver).
			Str("message", message),
	).Msg("msg sent")
	return nil
}
func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
