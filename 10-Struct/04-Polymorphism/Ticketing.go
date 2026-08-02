package main

import "fmt"

type BusTicket struct {
	Id                int
	PassengerName     string
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	BusType           string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}
type FlightTicket struct {
	Id               int
	PassengerName    string
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	AirplaneType     string
	PassportId       string
	PassengerType    string
	Price            int
}

func main() {
	busTicket := BusTicket{
		Id:                1,
		PassengerName:     "soheilsleep",
		DepartureCity:     "Tehran",
		ArrivalCity:       "Neyshabour",
		DepartureTime:     "12:00",
		BusType:           "VIP",
		DepartureTerminal: "Terminal 1",
		ArrivalTerminal:   "Terminal 2",
		NationalCode:      "1234567890",
		Price:             100,
	}
	flightTicket := FlightTicket{
		Id:               2,
		PassengerName:    "Erfan Mohseni",
		DepartureAirport: "Tehran",
		ArrivalAirport:   "London",
		DepartureTime:    "13:00",
		ArrivalTime:      "23:00",
		AirplaneType:     "Airbus",
		PassportId:       "0521436987",
		PassengerType:    "Adult",
		Price:            300,
	}
	printer := []TicketPrinter{busTicket, flightTicket}
	for _, printer := range printer {
		printer.PrintTicket()
	}
}

type TicketPrinter interface {
	PrintTicket()
}

func (Ticket BusTicket) PrintTicket() {
	fmt.Printf("BusTicket:\n Id:%d\n DepartureCity:%s\n ArrivalCity:%s\n DepartureTime:%s\n PassengerName:%s\n ", Ticket.Id, Ticket.DepartureCity, Ticket.ArrivalCity, Ticket.DepartureTime, Ticket.PassengerName)
	fmt.Printf("BusType:%s\n DepartureTerminal:%s\n ArrivalTerminal:%s\n NationalCode:%s\n Price:%d\n  ", Ticket.BusType, Ticket.DepartureTerminal, Ticket.ArrivalTerminal, Ticket.NationalCode, Ticket.Price)
}
func (Ticket FlightTicket) PrintTicket() {
	fmt.Printf("FlightTicket:\n Id:%d\n PassengerName:%s\n DepartureAirport:%s\n ArrivalAirport:%s\n  ", Ticket.Id, Ticket.PassengerName, Ticket.DepartureAirport, Ticket.ArrivalAirport)
	fmt.Printf("AirplaneType:%s\n DepartureTime:%s\n ArrivalTime:%s\n PassportId:%s\n PassengerType:%s\n  Price:%d\n  ", Ticket.AirplaneType, Ticket.DepartureTime, Ticket.ArrivalTime, Ticket.PassportId, Ticket.PassengerType, Ticket.Price)
}
