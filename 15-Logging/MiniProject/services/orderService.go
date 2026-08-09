package services

import (
	"errors"
	"fmt"
	"notification/ExternalServices"
	"notification/core"
	"notification/entities"
)

var logger = core.NewFileLogger()

type OrderService struct {
	ExternalServices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) (error, *entities.Order) {
	if !order.Status {

		return errors.New(fmt.Sprint("Order status is false")), nil
	}
	if order.Price < 150 {
		return errors.New(fmt.Sprint("Order price is not valid")), nil
	}
	fmt.Printf("Order Created: %v\n", order)
	logger.Info().Interface("order", order).Msgf("Order Created:")
	orderService.Notifier = ExternalServices.NewNotifier(order.NotificationType)
	logger.Info().Msgf("Notifier found:%T", orderService.Notifier)

	err := orderService.SendNotify(order.UserId, "Order Created")

	return err, order
}
func NewOrderService() *OrderService {
	return &OrderService{}
}
