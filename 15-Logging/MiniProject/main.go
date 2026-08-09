// main-> orderService -> emailService || smsService
package main

import (
	"notification/core"
	"notification/entities"
	"notification/services"
)

var logger = core.NewFileLogger()

func main() {
	order1 := entities.Order{
		Id:               1,
		UserFullName:     "soheil kazemi",
		UserId:           "09399850797",
		Price:            100,
		Status:           true,
		NotificationType: entities.Email,
	}
	order2 := entities.Order{
		Id:               2,
		UserFullName:     "erfan mohseni",
		UserId:           "09154857498",
		Price:            200,
		Status:           true,
		NotificationType: entities.Sms,
	}
	order3 := entities.Order{
		Id:               3,
		UserFullName:     "taha dehanvi",
		UserId:           "",
		Price:            150,
		Status:           true,
		NotificationType: entities.Nill,
	}
	orderService := services.NewOrderService()
	err, _ := orderService.CreateOrder(&order1)
	if err != nil {
		logger.Error().Interface("entityInfo", order1).Err(err).Msg("Error creating order")
	}
	err, _ = orderService.CreateOrder(&order2)
	if err != nil {
		logger.Error().Interface("entityInfo", order2).Err(err).Msg("Error creating order")
	}
	err, _ = orderService.CreateOrder(&order3)
	if err != nil {
		logger.Error().Interface("entityInfo", order3).Err(err).Msg("Error creating order")
	}
}
