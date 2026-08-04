package ExternalServices

import "fmt"

type NilNotifyService struct {
}

func (n NilNotifyService) SendNotify(receiver string, message string) {
	fmt.Printf("NilNotifyService(%s,%s)\n", receiver, message)
}
func NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}
