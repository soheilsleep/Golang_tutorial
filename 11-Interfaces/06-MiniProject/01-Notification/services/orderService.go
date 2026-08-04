package services

import (
	"fmt"
	"notification/ExternalServices"
	"notification/entities"
)

type OrderService struct {
	ExternalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	fmt.Printf("Order Created: %v\n", order)
	orderService.Notifier = ExternalServices.NewNotifier(order.NotificationType)
	orderService.SendNotify(order.UserId, "Order Created")

	return order
}
func NewOrderService() *OrderService {
	return &OrderService{}
}
