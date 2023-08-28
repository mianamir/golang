package abstract_factory

type Machine interface {
	Display() string
}

type CuttingMachine struct{}

func (m *CuttingMachine) Display() string {
	return "Cutting Machine"
}

type WeldingMachine struct{}

func (m *WeldingMachine) Display() string {
	return "Welding Machine"
}

type MachineFactory interface {
	CreateMachine() Machine
}

type CuttingMachineFactory struct{}

func (f *CuttingMachineFactory) CreateMachine() Machine {
	return &CuttingMachine{}
}

type WeldingMachineFactory struct{}

func (f *WeldingMachineFactory) CreateMachine() Machine {
	return &WeldingMachine{}
}
