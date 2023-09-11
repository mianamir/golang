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

// NewSingleton is a constructor function for creating a Singleton instance.
func NewSingleton() *Singleton {
	once.Do(func() {
		fmt.Println("Creating a new instance using sync.Once.")
		singletonInstance = &Singleton{value: 0}
	})
	return singletonInstance
}

func main() {
	// Using the constructor to create a Singleton instance
	singleton1 := NewSingleton()
	singleton2 := NewSingleton()

	// Accessing the values and checking if they are the same instance
	fmt.Printf("Value of singleton1: %d\n", singleton1.value)
	fmt.Printf("Value of singleton2: %d\n", singleton2.value)
	fmt.Printf("singleton1 is singleton2: %t\n", singleton1 == singleton2)
}
