# Online Movie Streaming Platform

Welcome to the Online Movie Streaming Platform project! This project demonstrates the implementation of SOLID principles in Golang to create an online movie streaming platform.

## Project Structure

The project is organized into modules, each focusing on different aspects of the online movie streaming platform:

- `movie_management.go`: Contains classes for managing movies and categories.
- `user_management.go`: Implements user account management and subscriptions.
- `billing.go`: Handles billing and payment processing.
- `interfaces.go`: Defines interfaces for different components.
- `main.go`: Entry point of the application.

## SOLID Principles Implementation

### Single Responsibility Principle (SRP)

The `Movie` and `Category` structs in `movie_management.go` each have a single responsibility. The `Movie` struct represents movie data, and the `Category` struct manages movie categories.

### Open/Closed Principle (OCP)

The `User` and `Subscription` interfaces in `user_management.go` follow the OCP. The `BasicSubscription` and `PremiumSubscription` types implement the `Subscription` interface, allowing for extension without modification.

### Liskov Substitution Principle (LSP)

The `Biller` interface in `billing.go` and the `BasicSubscription` and `PremiumSubscription` types adhere to LSP. Subtypes can be used interchangeably wherever the base type is expected.

### Interface Segregation Principle (ISP)

The `Playable` and `Billable` interfaces in `interfaces.go` are segregated into specific methods, preventing unnecessary method dependencies in implementing classes.

## Usage

1. Run the `main.go` script to see the simplified online movie streaming platform in action.

## Further Development

This project is a starting point to showcase SOLID principles in Golang. In real-world scenarios, you would expand on these concepts to create a complete and functional online movie streaming platform.

