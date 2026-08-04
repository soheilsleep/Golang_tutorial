package entities

type NotificationType string

const (
	Email NotificationType = "email"
	Sms   NotificationType = "sms"
	Nill  NotificationType = "nill"
)
