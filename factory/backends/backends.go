package backends

import (
	"fmt"
)

// PaymentBackend is an interface for payment backends.
type PaymentBackend interface {
	TokenizeCreditCard()
	ProcessPayment()
}

type credentials struct {
	username string
	key      string
}

// GetUsername returns the credentials username.
func (c *credentials) GetUsername() string {
	return c.username
}

// GetKey returns the credentials key.
func (c *credentials) GetKey() string {
	return c.key
}

// GetPaymentBackend is a factory for payments backends dependin on backend ID.
func GetPaymentBackend(paymentBackendID int) (PaymentBackend, error) {
	backends := map[int]PaymentBackend{
		1: &Cielo{
			credentials{
				username: "sample-username",
				key:      "sample-key",
			},
		},
		2: &Rede{
			credentials{
				username: "sample-username",
				key:      "sample-key",
			},
		},
	}

	if backend, exist := backends[paymentBackendID]; exist {
		return backend, nil
	}

	return nil, fmt.Errorf("invalid payment backend ID \"%d\"", paymentBackendID)
}
