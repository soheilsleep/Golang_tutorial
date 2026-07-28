package main

import "math/rand"

type CreditCard struct {
	CardNumber string
	ExpireDate string
	Cvv2       string
	BankName   string
}

func main() {
	cards := []CreditCard{
		{CardNumber: "5859831159530418", ExpireDate: "12/10", Cvv2: "123", BankName: "bank of soheilsleep"},
		{CardNumber: "8746851321651658", ExpireDate: "14/02", Cvv2: "321", BankName: "bank of soheilsleep"},
		{CardNumber: "7852365541356248", ExpireDate: "02/07", Cvv2: "415", BankName: "bank of soheilsleep"},
		{CardNumber: "5416531641656582", ExpireDate: "03/04", Cvv2: "468", BankName: "bank of soheilsleep"},
	}
	for _, card := range cards {
		if card.ExpireDate < "03/04" {
			println("your card is expired")
		}
		var remainAccount = getBankAccountRemainAmount(card.CardNumber, card.ExpireDate)
		println("remain amount of this card number", card.CardNumber, "is:", remainAccount)
	}
}
func getBankAccountRemainAmount(cardNumber string, expireDate string) int {
	min := 100000
	max := 1000000
	if expireDate < "03/04" {
		return 0

	}
	return (rand.Intn(max-min) + min)
}
