package main

import (
	"context"
	"fmt"
	"time"
)

// ResourceManager represents a resource that needs to be managed.
type ResourceManager struct {
	name string
}

// NewResourceManager creates a new ResourceManager instance.
func NewResourceManager(name string) *ResourceManager {
	return &ResourceManager{name: name}
}

// Close releases the resource when the context is canceled.
func (rm *ResourceManager) Close() {
	fmt.Printf("Closing resource: %s\n", rm.name)
	// Perform cleanup or resource release here.
}

func main() {
	// Create a context with a timeout of 3 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure the context is canceled when done.

	// Create and manage a resource within the context.
	rm := NewResourceManager("MyResource")

	select {
	case <-ctx.Done():
		// The context was canceled. Clean up the resource.
		rm.Close()
		fmt.Println("Resource management canceled.")
	case <-time.After(5 * time.Second):
		// Simulate some work with the resource.
		fmt.Println("Resource in use.")
	}

	// Rest of your program...
}
