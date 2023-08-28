package factorymethod

type Sensor interface {
	Measure() string
}

type TemperatureSensor struct{}

func (t *TemperatureSensor) Measure() string {
	return "Temperature: 25°C"
}

type PressureSensor struct{}

func (p *PressureSensor) Measure() string {
	return "Pressure: 100 kPa"
}

type SensorFactory interface {
	CreateSensor() Sensor
}

type TemperatureSensorFactory struct{}

func (f *TemperatureSensorFactory) CreateSensor() Sensor {
	return &TemperatureSensor{}
}

type PressureSensorFactory struct{}

func (f *PressureSensorFactory) CreateSensor() Sensor {
	return &PressureSensor{}
}
