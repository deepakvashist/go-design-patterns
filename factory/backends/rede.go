package backends

import (
	"fmt"
)

// Rede is a payment backend implementation.
type Rede struct {
	credentials
}

// TokenizeCreditCard tokenizes the credit card.
func (r *Rede) TokenizeCreditCard() {
	fmt.Println("REDE: Credit card tokenized")
}

// ProcessPayment processes the payment.
func (r *Rede) ProcessPayment() {
	fmt.Println("REDE: Payment finished successfully")
}
