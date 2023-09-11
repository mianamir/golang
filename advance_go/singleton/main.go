// Package main is the entry point of the Go program.
package main

import (
	"fmt"
	"sync"
)

// Singleton is a struct representing a singleton instance.
type Singleton struct {
	value int // value is the data stored in the Singleton instance.
}

var (
	singletonInstance *Singleton // singletonInstance stores the single instance of Singleton.
	once              sync.Once   // once is used to ensure initialization happens only once.
)

// GetSingletonInstance returns the singleton instance of Singleton.
func GetSingletonInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating a new instance using sync.Once.")
		singletonInstance = &Singleton{value: 0}
	})
	return singletonInstance
}

// SingletonUsingInit is a struct representing a singleton instance initialized using a method.
type SingletonUsingInit struct {
	value       int    // value is the data stored in the SingletonUsingInit instance.
	initialized bool   // initialized tracks whether the instance has been initialized.
}

var initSingletonInstance *SingletonUsingInit // initSingletonInstance stores the single instance of SingletonUsingInit.

// GetInitSingletonInstance returns the singleton instance of SingletonUsingInit, initializing it if necessary.
func GetInitSingletonInstance() *SingletonUsingInit {
	if initSingletonInstance == nil {
		fmt.Println("Creating a new instance using initialization.")
		initSingletonInstance = &SingletonUsingInit{value: 42, initialized: true}
	}
	return initSingletonInstance
}

func main() {
	// Using Approach 1: Using a struct with sync.Once
	singleton1 := GetSingletonInstance()
	singleton2 := GetSingletonInstance()

	// Using Approach 2: Using a struct with initialization
	initSingleton1 := GetInitSingletonInstance()
	initSingleton2 := GetInitSingletonInstance()

	// Accessing the values and checking if they are the same instance
	fmt.Printf("Value of singleton1: %d\n", singleton1.value)
	fmt.Printf("Value of singleton2: %d\n", singleton2.value)
	fmt.Printf("singleton1 is singleton2: %t\n", singleton1 == singleton2)

	fmt.Printf("Value of initSingleton1: %d\n", initSingleton1.value)
	fmt.Printf("Value of initSingleton2: %d\n", initSingleton2.value)
	fmt.Printf("initSingleton1 is initSingleton2: %t\n", initSingleton1 == initSingleton2)
}
