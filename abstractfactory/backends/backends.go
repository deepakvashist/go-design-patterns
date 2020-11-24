package backends

import (
	"fmt"
)

// PaymentBackend is an interface for payment backends.
type PaymentBackend interface {
	ProcessPayment()
}

// AntiFraudBackend is an interface for anti fraud backends.
type AntiFraudBackend interface {
	IsValidTransaction() bool
}

// Checkout is an interface for abstract checkout factory.
type Checkout interface {
	CreatePaymentBackend() PaymentBackend
	CreateAntiFraudBackend() AntiFraudBackend
}

// CheckoutSettingsA is a simple checkout implementation.
type CheckoutSettingsA struct{}

// CreatePaymentBackend returns a payment backend implementation.
func (c *CheckoutSettingsA) CreatePaymentBackend() PaymentBackend {
	return &Cielo{}
}

// CreateAntiFraudBackend returns a anti fraud backend implementation.
func (c *CheckoutSettingsA) CreateAntiFraudBackend() AntiFraudBackend {
	return &ClearSale{}
}

// CheckoutSettingsB is a simple checkout implementation.
type CheckoutSettingsB struct{}

// CreatePaymentBackend returns a payment backend implementation.
func (c *CheckoutSettingsB) CreatePaymentBackend() PaymentBackend {
	return &Rede{}
}

// CreateAntiFraudBackend returns a anti fraud backend implementation.
func (c *CheckoutSettingsB) CreateAntiFraudBackend() AntiFraudBackend {
	return &Serasa{}
}

// CheckoutFactory is a factory for checkout depending on checkout type.
func CheckoutFactory(checkoutType string) (Checkout, error) {
	checkouts := map[string]Checkout{
		"A": &CheckoutSettingsA{},
		"B": &CheckoutSettingsB{},
	}

	if checkout, exist := checkouts[checkoutType]; exist {
		return checkout, nil
	}

	return nil, fmt.Errorf("invalid checkout type \"%s\"", checkoutType)
}
