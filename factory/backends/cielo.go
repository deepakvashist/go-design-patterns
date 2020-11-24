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

// GetCredentialUsername get credentials username.
func (c *Cielo) GetCredentialUsername() string {
	return c.credentials.username
}

// GetCredentialKey get credentials key.
func (c *Cielo) GetCredentialKey() string {
	return c.credentials.key
}
