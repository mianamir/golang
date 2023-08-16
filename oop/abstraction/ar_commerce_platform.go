package main

import (
	"fmt"
)

// ARCommercePlatform defines the abstraction for buying products.
type ARCommercePlatform interface {
	BuyProduct(userID string, productID string, amount float64)
}

// VirtualFurniture represents a specific product type in the AR commerce platform.
type VirtualFurniture struct{}

// BuyProduct buys virtual furniture products.
func (vf *VirtualFurniture) BuyProduct(userID string, productID string, amount float64) {
	fmt.Printf("User %s bought virtual furniture %s for $%.2f.\n", userID, productID, amount)
}

// FashionItem represents another product type in the AR commerce platform.
type FashionItem struct{}

// BuyProduct buys fashion items.
func (fi *FashionItem) BuyProduct(userID string, productID string, amount float64) {
	fmt.Printf("User %s bought fashion item %s for $%.2f.\n", userID, productID, amount)
}

// ARCommerceFacade acts as a facade for interacting with the complex AR commerce platform.
type ARCommerceFacade struct {
	ARPlatform ARCommercePlatform
}

// BuyProduct provides a simplified interface for buying products.
func (acf *ARCommerceFacade) BuyProduct(userID string, productID string, amount float64) {
	acf.ARPlatform.BuyProduct(userID, productID, amount)
}

func main() {
	// Usage
	virtualFurniture := &VirtualFurniture{}
	arCommerceFacade := &ARCommerceFacade{ARPlatform: virtualFurniture}
	arCommerceFacade.BuyProduct("123", "VF123", 150)

	fashionItem := &FashionItem{}
	arCommerceFacade = &ARCommerceFacade{ARPlatform: fashionItem}
	arCommerceFacade.BuyProduct("456", "FI456", 75)
}
