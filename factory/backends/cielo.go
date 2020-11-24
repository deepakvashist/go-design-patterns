package backends

import (
	"fmt"
)

// Cielo is a payment backend implementation.
type Cielo struct {
	credentials
}

// TokenizeCreditCard tokenizes the credit card.
func (c *Cielo) TokenizeCreditCard() {
	fmt.Println("CIELO: Credit card tokenized")
}

// ProcessPayment processes the payment.
func (c *Cielo) ProcessPayment() {
	fmt.Println("CIELO: Payment finished successfully")
}
