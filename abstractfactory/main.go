package main

import (
	"github.com/deepakvashist/go-design-patterns/abstractfactory/backends"
)

func main() {
	checkoutA, _ := backends.CheckoutFactory("A")
	checkoutB, _ := backends.CheckoutFactory("B")

	paymentBackendA := checkoutA.CreatePaymentBackend()
	paymentBackendA.ProcessPayment()
	paymentBackendB := checkoutB.CreatePaymentBackend()
	paymentBackendB.ProcessPayment()
}
