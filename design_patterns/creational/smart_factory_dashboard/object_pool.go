package objectpool

type MachinePool struct {
	maxSize int
	pool    []Machine
}

func NewMachinePool(maxSize int) *MachinePool {
	return &MachinePool{
		maxSize: maxSize,
		pool:    []Machine{},
	}
}

func (p *MachinePool) GetMachine() Machine {
	if len(p.pool) < p.maxSize {
		newMachine := &MachineImpl{}
		p.pool = append(p.pool, newMachine)
		return newMachine
	} else {
		machine := p.pool[0]
		p.pool = p.pool[1:]
		return machine
	}
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
