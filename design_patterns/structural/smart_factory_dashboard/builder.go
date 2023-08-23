package builder

type ComplexMachineBuilder struct {
	machine Machine
}

func NewComplexMachineBuilder() *ComplexMachineBuilder {
	return &ComplexMachineBuilder{
		machine: &MachineImpl{},
	}
}

func (b *ComplexMachineBuilder) AddSensors(sensors []string) {
	b.machine.SetSensors(sensors)
}

func (b *ComplexMachineBuilder) AddName(name string) {
	b.machine.SetName(name)
}

func (b *ComplexMachineBuilder) GetMachine() Machine {
	return b.machine
}

type Machine interface {
	Display() string
	SetName(name string)
	SetSensors(sensors []string)
}

type MachineImpl struct {
	name    string
	sensors []string
}

func (m *MachineImpl) Display() string {
	return m.name
}

func (m *MachineImpl) SetName(name string) {
	m.name = name
}

func (m *MachineImpl) SetSensors(sensors []string) {
	m.sensors = sensors
}
