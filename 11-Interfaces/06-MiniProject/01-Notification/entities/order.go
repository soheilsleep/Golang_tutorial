package entities

type Order struct {
	Id               int
	UserFullName     string
	UserId           string
	Price            float64
	Status           bool
	NotificationType NotificationType
}