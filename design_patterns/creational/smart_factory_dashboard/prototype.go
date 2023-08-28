package prototype

import "github.com/mohit-kumar28/design-patterns-golang/builder"

type MachinePrototype interface {
	Clone() builder.Machine
}

type BasicMachine struct{}

func (b *BasicMachine) Clone() builder.Machine {
	return &builder.MachineImpl{
		name: "Basic Machine",
	}
}
