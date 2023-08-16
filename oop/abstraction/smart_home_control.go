package main

import (
	"fmt"
)

// SmartHomeDevice defines the abstraction for executing device control requests.
type SmartHomeDevice interface {
	Execute()
}

// ThermostatControl encapsulates thermostat control requests.
type ThermostatControl struct {
	DeviceID string
	Action   string
}

// Execute executes thermostat control actions.
func (tc *ThermostatControl) Execute() {
	fmt.Printf("Controlled thermostat %s: %s.\n", tc.DeviceID, tc.Action)
}

// LightingControl encapsulates lighting control requests.
type LightingControl struct {
	DeviceID string
	Action   string
}

// Execute executes lighting control actions.
func (lc *LightingControl) Execute() {
	fmt.Printf("Controlled lighting %s: %s.\n", lc.DeviceID, lc.Action)
}

// SmartHomeControl holds a collection of device control commands and executes them on demand.
type SmartHomeControl struct {
	Commands []SmartHomeDevice
}

// AddCommand adds a device control command.
func (shc *SmartHomeControl) AddCommand(command SmartHomeDevice) {
	shc.Commands = append(shc.Commands, command)
}

// ExecuteCommands executes all the device control commands.
func (shc *SmartHomeControl) ExecuteCommands() {
	for _, command := range shc.Commands {
		command.Execute()
	}
}

func main() {
	// Usage
	thermostatControl := &ThermostatControl{DeviceID: "T123", Action: "Increase temperature"}
	lightingControl := &LightingControl{DeviceID: "L456", Action: "Turn off"}

	smartHomeControl := &SmartHomeControl{}
	smartHomeControl.AddCommand(thermostatControl)
	smartHomeControl.AddCommand(lightingControl)

	smartHomeControl.ExecuteCommands()
}
