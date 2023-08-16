package main

import (
	"fmt"
)

// SmartFactoryControl defines the abstraction for performing maintenance.
type SmartFactoryControl interface {
	PerformMaintenance(equipmentID string, amount float64)
}

// MaintenanceObserver defines the observer interface for maintenance events.
type MaintenanceObserver interface {
	NotifyMaintenance(equipmentID string)
}

// RobotMaintenance represents robot maintenance equipment.
type RobotMaintenance struct{}

// PerformMaintenance performs maintenance on a robot.
func (rm *RobotMaintenance) PerformMaintenance(equipmentID string, amount float64) {
	fmt.Printf("Performed maintenance on robot %s for $%.2f.\n", equipmentID, amount)
	rm.NotifyMaintenance(equipmentID)
}

// NotifyMaintenance notifies maintenance events for a robot.
func (rm *RobotMaintenance) NotifyMaintenance(equipmentID string) {
	fmt.Printf("Robot %s maintenance event notified.\n", equipmentID)
}

// ConveyorBeltMaintenance represents conveyor belt maintenance equipment.
type ConveyorBeltMaintenance struct{}

// PerformMaintenance performs maintenance on a conveyor belt.
func (cbm *ConveyorBeltMaintenance) PerformMaintenance(equipmentID string, amount float64) {
	fmt.Printf("Performed maintenance on conveyor belt %s for $%.2f.\n", equipmentID, amount)
	cbm.NotifyMaintenance(equipmentID)
}

// NotifyMaintenance notifies maintenance events for a conveyor belt.
func (cbm *ConveyorBeltMaintenance) NotifyMaintenance(equipmentID string) {
	fmt.Printf("Conveyor belt %s maintenance event notified.\n", equipmentID)
}

func main() {
	// Usage
	robotMaintenance := &RobotMaintenance{}
	conveyorBeltMaintenance := &ConveyorBeltMaintenance{}

	robotMaintenance.PerformMaintenance("R123", 500)
	conveyorBeltMaintenance.PerformMaintenance("CB456", 300)
}
