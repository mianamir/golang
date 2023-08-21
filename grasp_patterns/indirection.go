package grasp_patterns

import (
	"fmt"
)

// PaymentProcessor defines the interface for payment processing
type PaymentProcessor interface {
	ProcessPayment(amount float64) bool
}

// OrderModule handles order-related operations
type OrderModule struct {
	paymentProcessor PaymentProcessor
}

type Order struct {
	OrderID int
	Total   float64
}

func (om *OrderModule) PlaceOrder(order Order) bool {
	return om.paymentProcessor.ProcessPayment(order.Total)
}

// PayPalProcessor implements PaymentProcessor
type PayPalProcessor struct {
	// ...
}

func (pp *PayPalProcessor) ProcessPayment(amount float64) bool {
	// Logic to process payment via PayPal
	fmt.Printf("Payment of $%.2f processed via PayPal\n", amount)
	return true
}

// CreditCardProcessor implements PaymentProcessor
type CreditCardProcessor struct {
	// ...
}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) bool {
	// Logic to process payment via credit card
	fmt.Printf("Payment of $%.2f processed via Credit Card\n", amount)
	return true
}

func main() {
	paypalProcessor := &PayPalProcessor{}
	orderModule := OrderModule{paymentProcessor: paypalProcessor}

	order := Order{OrderID: 123, Total: 50.75}
	success := orderModule.PlaceOrder(order)

	if success {
		fmt.Println("Order placed successfully!")
	} else {
		fmt.Println("Order placement failed.")
	}
}
