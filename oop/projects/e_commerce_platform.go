package main

import "fmt"

// Product is the interface defining the contract for obtaining product details.
type Product interface {
	GetDetails() string
}

// Electronics represents an electronics product.
type Electronics struct {
	name  string
	price float64
}

// GetDetails returns details about the electronics product.
func (e *Electronics) GetDetails() string {
	return fmt.Sprintf("Electronics: %s, Price: $%.2f", e.name, e.price)
}

// Clothing represents a clothing product.
type Clothing struct {
	name  string
	price float64
}

// GetDetails returns details about the clothing product.
func (c *Clothing) GetDetails() string {
	return fmt.Sprintf("Clothing: %s, Price: $%.2f", c.name, c.price)
}

// PaymentGateway is the interface defining the contract for processing payments.
type PaymentGateway interface {
	ProcessPayment(amount float64) bool
}

// CreditCardPayment represents a credit card payment gateway.
type CreditCardPayment struct{}

// ProcessPayment simulates processing a credit card payment.
func (cc *CreditCardPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processed $%.2f via credit card.\n", amount)
	return true
}

// PayPalPayment represents a PayPal payment gateway.
type PayPalPayment struct{}

// ProcessPayment simulates processing a PayPal payment.
func (pp *PayPalPayment) ProcessPayment(amount float64) bool {
	fmt.Printf("Processed $%.2f via PayPal.\n", amount)
	return true
}

// OrderProcessor manages order processing.
type OrderProcessor struct {
	paymentGateway PaymentGateway
}

// NewOrderProcessor creates a new OrderProcessor instance.
func NewOrderProcessor(paymentGateway PaymentGateway) *OrderProcessor {
	return &OrderProcessor{paymentGateway: paymentGateway}
}

// ProcessOrder processes an order and performs payment.
func (op *OrderProcessor) ProcessOrder(product Product, quantity int, totalAmount float64) bool {
	orderDetails := product.GetDetails()
	fmt.Printf("Processing order for: %s, Quantity: %d\n", orderDetails, quantity)

	// Process payment
	if op.paymentGateway.ProcessPayment(totalAmount) {
		fmt.Println("Order successfully processed.")
		return true
	} else {
		fmt.Println("Payment processing failed. Order not processed.")
		return false
	}
}

// InventoryObserver is the interface defining the contract for updating inventory.
type InventoryObserver interface {
	UpdateInventory(product Product, quantity int)
}

// InventoryManager manages inventory updates.
type InventoryManager struct{}

// UpdateInventory updates the inventory for a product and quantity.
func (im *InventoryManager) UpdateInventory(product Product, quantity int) {
	fmt.Printf("Updated inventory for %s, Quantity: %d\n", product.GetDetails(), quantity)
}

// ECommercePlatform represents the e-commerce platform.
type ECommercePlatform struct {
	orderProcessor    *OrderProcessor
	inventoryObserver InventoryObserver
}

// NewECommercePlatform creates a new ECommercePlatform instance.
func NewECommercePlatform(orderProcessor *OrderProcessor, inventoryObserver InventoryObserver) *ECommercePlatform {
	return &ECommercePlatform{
		orderProcessor:    orderProcessor,
		inventoryObserver: inventoryObserver,
	}
}

// ProcessOrder processes an order, updates inventory, and performs payment.
func (ecp *ECommercePlatform) ProcessOrder(product Product, quantity int, totalAmount float64) {
	if ecp.orderProcessor.ProcessOrder(product, quantity, totalAmount) {
		ecp.inventoryObserver.UpdateInventory(product, quantity)
	}
}

func main() {
	// Creating products
	electronics := &Electronics{name: "Smartphone", price: 500}
	clothing := &Clothing{name: "T-Shirt", price: 20}

	// Creating payment gateways
	creditCardPayment := &CreditCardPayment{}
	payPalPayment := &PayPalPayment{}

	// Creating inventory manager
	inventoryManager := &InventoryManager{}

	// Creating the E-commerce platform
	orderProcessor := NewOrderProcessor(creditCardPayment)
	eCommercePlatform := NewECommercePlatform(orderProcessor, inventoryManager)

	// Processing orders
	eCommercePlatform.ProcessOrder(electronics, 2, 1000)
}
