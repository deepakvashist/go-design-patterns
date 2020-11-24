package backends

import (
	"fmt"
)

// Rede is a payment backend implementation.
type Rede struct{}

// ProcessPayment processes the payment.
func (r *Rede) ProcessPayment() {
	fmt.Println("REDE: Payment finished successfully")
}
