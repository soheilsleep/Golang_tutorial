// main-> orderService -> emailService || smsService
package main

import (
	"notification/entities"
	"notification/services"
)

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
		UserId:           "09561484562",
		Price:            150,
		Status:           true,
		NotificationType: entities.Nill,
	}
	orderService := services.NewOrderService()
	orderService.CreateOrder(&order1)
	orderService.CreateOrder(&order2)
	orderService.CreateOrder(&order3)
}
