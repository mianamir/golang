package projects

import (
	"fmt"
)

// Machine represents a basic machine with a name.
type Machine struct {
	Name string
}

// Operate defines the operation of a machine.
func (m *Machine) Operate() {
	// This method is intentionally left empty in the base class.
}

// ConveyorBelt represents a type of machine - a conveyor belt.
type ConveyorBelt struct {
	Machine
}

// Operate performs the operation of a conveyor belt.
func (cb *ConveyorBelt) Operate() {
	fmt.Printf("%s conveyor belt is moving.\n", cb.Name)
}

// RobotArm represents a type of machine - a robot arm.
type RobotArm struct {
	Machine
}

// Operate performs the operation of a robot arm.
func (ra *RobotArm) Operate() {
	fmt.Printf("%s robot arm is assembling.\n", ra.Name)
}

// ProductionLine represents a production line with multiple machines.
type ProductionLine struct {
	Machines []Machine
}

// Run executes the production process on the production line.
func (pl *ProductionLine) Run() {
	for _, machine := range pl.Machines {
		machine.Operate()
	}
}

func main() {
	// Create instances of ConveyorBelt and RobotArm
	conveyorBelt := ConveyorBelt{Machine: Machine{Name: "Main"}}
	robotArm := RobotArm{Machine: Machine{Name: "Assembly"}}

	// Create a production line with machines
	productionLine := ProductionLine{
		Machines: []Machine{&conveyorBelt, &robotArm},
	}

	// Use the Production Line
	productionLine.Run()
}
