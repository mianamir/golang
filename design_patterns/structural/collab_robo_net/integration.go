package main

import "fmt"

// Adapter Pattern: SystemIntegration
// This pattern allows objects with incompatible interfaces to collaborate.
type ExternalSystem interface {
	SendData(data string)
	ReceiveData() string
}

type SystemIntegration struct {
	externalSystem ExternalSystem
}

func NewSystemIntegration(externalSystem ExternalSystem) *SystemIntegration {
	return &SystemIntegration{externalSystem: externalSystem}
}

func (si *SystemIntegration) SendDataToExternalSystem(data string) {
	// Implementation for sending data to external factory systems
	si.externalSystem.SendData(data)
}

func (si *SystemIntegration) ReceiveDataFromExternalSystem() string {
	// Implementation for receiving data from external systems
	return si.externalSystem.ReceiveData()
}

// Example ExternalSystem Implementation
type ExampleExternalSystem struct{}

func (ees *ExampleExternalSystem) SendData(data string) {
	fmt.Printf("Sending data to external system: %s\n", data)
}

func (ees *ExampleExternalSystem) ReceiveData() string {
	fmt.Println("Receiving data from external system")
	return "Received data"
}

func main() {
	// Example usage
	externalSystem := &ExampleExternalSystem{}

	systemIntegration := NewSystemIntegration(externalSystem)

	systemIntegration.SendDataToExternalSystem("Data to be sent")

	receivedData := systemIntegration.ReceiveDataFromExternalSystem()
	fmt.Println("Received:", receivedData)
}
