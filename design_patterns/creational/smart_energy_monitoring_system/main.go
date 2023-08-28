package main

import "fmt"

// EnergySensor interface
type EnergySensor interface {
	MeasureEnergy() string
}

// VoltageSensor struct
type VoltageSensor struct{}

func (vs *VoltageSensor) MeasureEnergy() string {
	return "Voltage: 220V"
}

// CurrentSensor struct
type CurrentSensor struct{}

func (cs *CurrentSensor) MeasureEnergy() string {
	return "Current: 5A"
}

// EnergySensorFactory interface
type EnergySensorFactory interface {
	CreateSensor() EnergySensor
}

// VoltageSensorFactory struct
type VoltageSensorFactory struct{}

func (vsf *VoltageSensorFactory) CreateSensor() EnergySensor {
	return &VoltageSensor{}
}

// CurrentSensorFactory struct
type CurrentSensorFactory struct{}

func (csf *CurrentSensorFactory) CreateSensor() EnergySensor {
	return &CurrentSensor{}
}

// EnergyDashboard struct
type EnergyDashboard struct {
	sensors []string
}

func NewEnergyDashboard() *EnergyDashboard {
	return &EnergyDashboard{
		sensors: make([]string, 0),
	}
}

func (ed *EnergyDashboard) AddSensor(sensor string) {
	ed.sensors = append(ed.sensors, sensor)
}

func (ed *EnergyDashboard) GetSensors() []string {
	return ed.sensors
}

func main() {
	// Factory Pattern
	voltageSensorFactory := &VoltageSensorFactory{}
	currentSensorFactory := &CurrentSensorFactory{}

	voltageSensor := voltageSensorFactory.CreateSensor()
	currentSensor := currentSensorFactory.CreateSensor()

	// Singleton Pattern
	dashboard := NewEnergyDashboard()
	dashboard.AddSensor(voltageSensor.MeasureEnergy())
	dashboard.AddSensor(currentSensor.MeasureEnergy())

	fmt.Println("Energy Consumption Dashboard")
	fmt.Println("------------------------------")
	for _, sensor := range dashboard.GetSensors() {
		fmt.Println(sensor)
	}
}
