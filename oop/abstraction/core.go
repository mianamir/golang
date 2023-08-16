package main

import (
	"fmt"
)

// PaymentGateway defines the abstract interface for payment processing.
type PaymentGateway interface {
	ProcessPayment(amount float64)
}

// CreditCardPayment represents a payment processing method via credit card.
type CreditCardPayment struct{}

// ProcessPayment processes a payment via credit card.
func (cc *CreditCardPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processed $%.2f via credit card.\n", amount)
}

// PaypalPayment represents a payment processing method via PayPal.
type PaypalPayment struct{}

// ProcessPayment processes a payment via PayPal.
func (pp *PaypalPayment) ProcessPayment(amount float64) {
	fmt.Printf("Processed $%.2f via PayPal.\n", amount)
}

func main() {
	// Usage
	var payment PaymentGateway

	payment = &CreditCardPayment{}
	payment.ProcessPayment(100.0)

	payment = &PaypalPayment{}
	payment.ProcessPayment(50.0)
}
