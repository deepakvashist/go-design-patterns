package main

import (
	"github.com/deepakvashist/go-design-patterns/factory/backends"
)

func main() {
	cielo, _ := backends.GetPaymentBackend(1)
	rede, _ := backends.GetPaymentBackend(2)

	cielo.TokenizeCreditCard()
	cielo.ProcessPayment()
	rede.TokenizeCreditCard()
	rede.ProcessPayment()
}
