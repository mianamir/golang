package main

import (
	"fmt"
)

// Define the Customer struct
type Customer struct {
	Name   string
	Email  string
	Orders []string
}

// Define a method to place an order for a customer
func (c *Customer) PlaceOrder(order string) {
	c.Orders = append(c.Orders, order)
	fmt.Printf("%s placed an order for %s\n", c.Name, order)
}

func main() {
	// Create instances of the Customer struct
	customer1 := Customer{Name: "Muhammad", Email: "muhammad@example.com"}
	customer2 := Customer{Name: "Amir", Email: "amir@example.com"}

	// Call methods on the objects
	customer1.PlaceOrder("Shoes")
	customer2.PlaceOrder("Phone")
}
