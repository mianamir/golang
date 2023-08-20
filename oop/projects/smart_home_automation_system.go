package main

import "fmt"

// SmartHomeDevice is the interface defining the contract for smart home devices.
type SmartHomeDevice interface {
	PerformAction()
}

// Light represents a light smart home device.
type Light struct{}

// PerformAction turns on the light.
func (l *Light) PerformAction() {
	fmt.Println("Light is turned on.")
}

// Thermostat represents a thermostat smart home device.
type Thermostat struct{}

// PerformAction adjusts the thermostat temperature.
func (t *Thermostat) PerformAction() {
	fmt.Println("Thermostat temperature is adjusted.")
}

// Speaker represents a speaker smart home device.
type Speaker struct{}

// PerformAction plays music on the speaker.
func (s *Speaker) PerformAction() {
	fmt.Println("Speaker is playing music.")
}

// SmartHomeControl is the interface defining the contract for smart home control.
type SmartHomeControl interface {
	Execute()
}

// ControlCommand represents a control command for smart home devices.
type ControlCommand struct {
	device SmartHomeDevice
}

// Execute performs the action on the associated smart home device.
func (cc *ControlCommand) Execute() {
	cc.device.PerformAction()
}

// SmartHomeControlSystem manages smart home control commands.
type SmartHomeControlSystem struct {
	commands []SmartHomeControl
}

// AddCommand adds a smart home control command to the system.
func (shcs *SmartHomeControlSystem) AddCommand(command SmartHomeControl) {
	shcs.commands = append(shcs.commands, command)
}

// ExecuteCommands executes all smart home control commands.
func (shcs *SmartHomeControlSystem) ExecuteCommands() {
	for _, command := range shcs.commands {
		command.Execute()
	}
}

// MaintenanceNotifier notifies maintenance events to observers.
type MaintenanceNotifier struct {
	observers []MaintenanceObserver
}

// AddObserver adds a maintenance observer to the notifier.
func (mn *MaintenanceNotifier) AddObserver(observer MaintenanceObserver) {
	mn.observers = append(mn.observers, observer)
}

// RemoveObserver removes a maintenance observer from the notifier.
func (mn *MaintenanceNotifier) RemoveObserver(observer MaintenanceObserver) {
	for i, obs := range mn.observers {
		if obs == observer {
			mn.observers = append(mn.observers[:i], mn.observers[i+1:]...)
			break
		}
	}
}

// NotifyMaintenance notifies a maintenance event to all observers.
func (mn *MaintenanceNotifier) NotifyMaintenance(deviceName string) {
	for _, observer := range mn.observers {
		observer.NotifyMaintenance(deviceName)
	}
}

// MaintenanceObserver is the interface defining the contract for maintenance observers.
type MaintenanceObserver interface {
	NotifyMaintenance(deviceName string)
}

// LightMaintenanceObserver observes maintenance events for lights.
type LightMaintenanceObserver struct{}

// NotifyMaintenance notifies a light maintenance event.
func (lmo *LightMaintenanceObserver) NotifyMaintenance(deviceName string) {
	fmt.Printf("Maintenance alert for %s: Check light.\n", deviceName)
}

// ThermostatMaintenanceObserver observes maintenance events for thermostats.
type ThermostatMaintenanceObserver struct{}

// NotifyMaintenance notifies a thermostat maintenance event.
func (tmo *ThermostatMaintenanceObserver) NotifyMaintenance(deviceName string) {
	fmt.Printf("Maintenance alert for %s: Check thermostat.\n", deviceName)
}

// SpeakerMaintenanceObserver observes maintenance events for speakers.
type SpeakerMaintenanceObserver struct{}

// NotifyMaintenance notifies a speaker maintenance event.
func (smo *SpeakerMaintenanceObserver) NotifyMaintenance(deviceName string) {
	fmt.Printf("Maintenance alert for %s: Check speaker.\n", deviceName)
}

func main() {
	// Creating the smart home control system
	smartHomeControlSystem := SmartHomeControlSystem{}

	// Creating specific smart home devices and commands
	light := &Light{}
	thermostat := &Thermostat{}
	speaker := &Speaker{}

	lightCommand := &ControlCommand{device: light}
	thermostatCommand := &ControlCommand{device: thermostat}
	speakerCommand := &ControlCommand{device: speaker}

	smartHomeControlSystem.AddCommand(lightCommand)
	smartHomeControlSystem.AddCommand(thermostatCommand)
	smartHomeControlSystem.AddCommand(speakerCommand)

	// Executing commands
	smartHomeControlSystem.ExecuteCommands()

	// Creating maintenance notifier and observers
	maintenanceNotifier := MaintenanceNotifier{}
	lightObserver := &LightMaintenanceObserver{}
	thermostatObserver := &ThermostatMaintenanceObserver{}
	speakerObserver := &SpeakerMaintenanceObserver{}

	// Adding observers to the notifier
	maintenanceNotifier.AddObserver(lightObserver)
	maintenanceNotifier.AddObserver(thermostatObserver)
	maintenanceNotifier.AddObserver(speakerObserver)

	// Notifying maintenance events
	maintenanceNotifier.NotifyMaintenance("Light")
	maintenanceNotifier.NotifyMaintenance("Thermostat")
	maintenanceNotifier.NotifyMaintenance("Speaker")
}
