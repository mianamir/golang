package main

import (
	"fmt"
)

// Product is a struct representing a product
type Product struct {
    name  string
    price int
}

// NewProduct creates a new Product instance
func NewProduct(name string, price int) *Product {
    return &Product{name: name, price: price}
}

// Price returns the price of the product
func (p *Product) Price() int {
    return p.price
}

// SetPrice sets the price of the product with validation
func (p *Product) SetPrice(newPrice int) error {
    if newPrice < 0 {
        return fmt.Errorf("Price cannot be negative")
    }
    p.price = newPrice
    fmt.Printf("Price updated for %s: $%d\n", p.name, newPrice)
    return nil
}

// DeletePrice deletes the price of the product
func (p *Product) DeletePrice() {
    fmt.Printf("Price deleted for %s\n", p.name)
    p.price = 0
}

// CreateProduct creates a new Product instance
func CreateProduct(name string, price int) *Product {
    return NewProduct(name, price)
}

// ValidatePositivePrice is a decorator to validate positive prices
func ValidatePositivePrice(setPriceFunc func(*Product, int) error) func(*Product, int) error {
    return func(p *Product, price int) error {
        if price < 0 {
            return fmt.Errorf("Price cannot be negative")
        }
        return setPriceFunc(p, price)
    }
}

// CacheResult is a decorator to cache the results of a function
func CacheResult(fn func(*Product, int) error) func(*Product, int) error {
    cache := make(map[string]error)

    return func(p *Product, price int) error {
        key := fmt.Sprintf("%s_%d", fn, price)
        if cachedResult, ok := cache[key]; ok {
            return cachedResult
        }
        result := fn(p, price)
        cache[key] = result
        return result
    }
}

// RetryOnFailure is a decorator to retry a function on failure
func RetryOnFailure(maxRetries int, fn func(*Product, int) error) func(*Product, int) error {
    return func(p *Product, price int) error {
        for i := 0; i < maxRetries; i++ {
            if err := fn(p, price); err == nil {
                return nil
            } else {
                fmt.Printf("Retry failed: %v\n", err)
            }
        }
        return fmt.Errorf("Max retries (%d) exceeded", maxRetries)
    }
}

func main() {
    // Using the enhanced Product struct
    product1 := CreateProduct("Laptop", 1000)

    // Apply decorators to set the price
    setPriceWithValidation := ValidatePositivePrice(product1.SetPrice)
    setPriceWithCache := CacheResult(setPriceWithValidation)
    setPriceWithRetry := RetryOnFailure(3, setPriceWithCache)

    // Using the enhanced SetPrice method
    if err := setPriceWithRetry(product1, 1200); err == nil {
        // This will log the price change and cache the result
    } else {
        fmt.Printf("Error: %v\n", err)
    }

    if err := setPriceWithRetry(product1, -200); err != nil {
        fmt.Printf("Error: %v\n", err)
        // This will retry and eventually raise an error
    }
}
