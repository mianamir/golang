# Singleton Design Pattern in Go

This repository contains a Go implementation of the Singleton design pattern using two different approaches. The Singleton pattern ensures that a class has only one instance and provides a global point of access to that instance.

## Code Overview

### Singleton Structs

- `Singleton`: A struct representing a singleton instance with a `value` field, which can store an integer value.

- `SingletonUsingInit`: A struct representing a singleton instance initialized using a method. It includes a `value` field and an optional `initialized` field.

### Singleton Instances

- `singletonInstance`: A global variable to hold the single instance of the `Singleton` struct.

- `initSingletonInstance`: A global variable to hold the single instance of the `SingletonUsingInit` struct.

### `sync.Once`

The code uses the `sync.Once` construct to ensure that initialization code is executed only once, providing a safe mechanism for creating singletons.

### Functionality

- `GetSingletonInstance()`: Returns the singleton instance of the `Singleton` struct. It ensures that the instance is created only once using the `sync.Once` construct.

- `GetInitSingletonInstance()`: Returns the singleton instance of the `SingletonUsingInit` struct. It checks whether the instance exists and initializes it with a value of 42 if it doesn't.

### Usage in Main Function

In the `main` function, the code demonstrates the use of both approaches:

- Calls `GetSingletonInstance()` twice to retrieve the same `Singleton` instance, showcasing that the `sync.Once` ensures a single instance.

- Calls `GetInitSingletonInstance()` twice to retrieve the same `SingletonUsingInit` instance, demonstrating that it initializes only once.

The code also prints messages to the console to indicate whether a new instance is created or an existing instance is reused. It displays the values of the singleton instances and checks if they are the same instance.

## Running the Code

To run the code, simply execute the `main.go` file. You should see the output demonstrating the Singleton pattern in action.

```bash
go run main.go
