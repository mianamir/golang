package main

import (
	"fmt"
)

// VRExperiencePlatform defines the abstraction for purchasing experiences.
type VRExperiencePlatform interface {
	PurchaseExperience(userID string, experienceID string, amount float64)
}

// BaseExperience wraps concrete experiences, acting as a base for decorators.
type BaseExperience struct {
	WrappedExperience VRExperiencePlatform
}

// PurchaseExperience delegates the purchase to the wrapped experience.
func (be *BaseExperience) PurchaseExperience(userID string, experienceID string, amount float64) {
	be.WrappedExperience.PurchaseExperience(userID, experienceID, amount)
}

// GameExperience represents a game experience.
type GameExperience struct{}

// PurchaseExperience purchases a game experience.
func (ge *GameExperience) PurchaseExperience(userID string, experienceID string, amount float64) {
	fmt.Printf("User %s purchased game experience %s for $%.2f.\n", userID, experienceID, amount)
}

// TourExperience represents a tour experience.
type TourExperience struct{}

// PurchaseExperience purchases a tour experience.
func (te *TourExperience) PurchaseExperience(userID string, experienceID string, amount float64) {
	fmt.Printf("User %s purchased tour experience %s for $%.2f.\n", userID, experienceID, amount)
}

func main() {
	// Usage
	gameExperience := &GameExperience{}
	decoratedGameExperience := &BaseExperience{WrappedExperience: gameExperience}

	decoratedGameExperience.PurchaseExperience("123", "VR123", 20)

	tourExperience := &TourExperience{}
	decoratedTourExperience := &BaseExperience{WrappedExperience: tourExperience}

	decoratedTourExperience.PurchaseExperience("124", "VR456", 30)
}
