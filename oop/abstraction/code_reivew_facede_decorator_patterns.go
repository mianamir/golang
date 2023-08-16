package main

import (
	"fmt"
)

// Distinct Use Cases:

/*
In the following code review, we'll examine the use of the Decorator and Facade patterns
in the provided examples: "Virtual Reality - VR Experience Platform" (Decorator Pattern)
and "AR Commerce Platform" (Facade Pattern). These patterns serve different purposes and are
applied in distinct contexts. Let's analyze the reasons behind their usage and their respective applications.
*/

// Virtual Reality - VR Experience Platform (Decorator Pattern):

/*
The Decorator pattern is employed in the "Virtual Reality - VR Experience Platform" example
to dynamically add responsibilities (features or behaviors) to objects without altering their structure.
This pattern allows for flexible extension of objects' functionality.
*/

// VRExperiencePlatform defines the abstraction for purchasing experiences.
type VRExperiencePlatform interface {
	BuyProduct(userID string, productID string, amount float64)
}

// BaseExperience wraps concrete experiences, acting as a base for decorators.
type BaseExperience struct {
	WrappedExperience VRExperiencePlatform
}

// BuyProduct delegates the purchase to the wrapped experience.
func (be *BaseExperience) BuyProduct(userID string, productID string, amount float64) {
	be.WrappedExperience.BuyProduct(userID, productID, amount)
}

// GameExperience represents a specific product type in the VR experience platform.
type GameExperience struct{}

// BuyProduct purchases a game experience.
func (ge *GameExperience) BuyProduct(userID string, productID string, amount float64) {
	fmt.Printf("User %s bought game experience %s for $%.2f.\n", userID, productID, amount)
}

// TourExperience represents another product type in the VR experience platform.
type TourExperience struct{}

// BuyProduct purchases a tour experience.
func (te *TourExperience) BuyProduct(userID string, productID string, amount float64) {
	fmt.Printf("User %s bought tour experience %s for $%.2f.\n", userID, productID, amount)
}

// AR Commerce Platform (Facade Pattern):

/*
The Facade pattern is applied in the "AR Commerce Platform" example to provide a simplified,
higher-level interface to a complex subsystem. It shields clients from the underlying complexities
of the AR commerce platform.
*/

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
	// Virtual Reality - VR Experience Platform (Decorator Pattern) Usage
	gameExperience := &GameExperience{}
	decoratedGameExperience := &BaseExperience{WrappedExperience: gameExperience}
	decoratedGameExperience.BuyProduct("123", "GE123", 20)

	tourExperience := &TourExperience{}
	decoratedTourExperience := &BaseExperience{WrappedExperience: tourExperience}
	decoratedTourExperience.BuyProduct("456", "TE456", 30)

	// AR Commerce Platform (Facade Pattern) Usage
	virtualFurniture := &VirtualFurniture{}
	arCommerceFacade := &ARCommerceFacade{ARPlatform: virtualFurniture}
	arCommerceFacade.BuyProduct("789", "VF789", 150)

	fashionItem := &FashionItem{}
	arCommerceFacade = &ARCommerceFacade{ARPlatform: fashionItem}
	arCommerceFacade.BuyProduct("101", "FI101", 75)
}
