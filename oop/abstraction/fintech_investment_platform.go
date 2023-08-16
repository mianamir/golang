package main

import (
	"fmt"
)

// InvestmentPlatform defines the abstract interface for investments.
type InvestmentPlatform interface {
	Invest(userID string, amount float64)
}

// StockInvestment represents an investment type in stocks.
type StockInvestment struct{}

// Invest implements the investment in stocks.
func (si *StockInvestment) Invest(userID string, amount float64) {
	fmt.Printf("User %s invested $%.2f in stocks.\n", userID, amount)
}

// RealEstateInvestment represents an investment type in real estate.
type RealEstateInvestment struct{}

// Invest implements the investment in real estate.
func (re *RealEstateInvestment) Invest(userID string, amount float64) {
	fmt.Printf("User %s invested $%.2f in real estate.\n", userID, amount)
}

// BitcoinInvestment represents an investment type in bitcoin.
type BitcoinInvestment struct{}

// Invest implements the investment in bitcoin.
func (bi *BitcoinInvestment) Invest(userID string, amount float64) {
	fmt.Printf("User %s invested $%.2f in bitcoin.\n", userID, amount)
}

// InvestmentFactory represents a factory for creating investment instances.
type InvestmentFactory struct{}

// CreateInvestment creates an investment instance based on the investment type.
func (ifc *InvestmentFactory) CreateInvestment(investmentType string) InvestmentPlatform {
	switch investmentType {
	case "stock":
		return &StockInvestment{}
	case "real_estate":
		return &RealEstateInvestment{}
	case "bitcoin":
		return &BitcoinInvestment{}
	default:
		return nil
	}
}

func main() {
	// Usage
	investmentFactory := InvestmentFactory{}

	// Creating and using different investment types
	stockInvestment := investmentFactory.CreateInvestment("stock")
	stockInvestment.Invest("123", 1000)

	realEstateInvestment := investmentFactory.CreateInvestment("real_estate")
	realEstateInvestment.Invest("124", 5000)

	bitcoinInvestment := investmentFactory.CreateInvestment("bitcoin")
	bitcoinInvestment.Invest("125", 10000)
}
