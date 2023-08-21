package grasp_patterns

import (
	"fmt"
)

// SensorService defines the interface for interacting with sensors
type SensorService interface {
	ReadSensorValue(sensorID int) float64
}

// TemperatureSensor implements SensorService for temperature measurements
type TemperatureSensor struct {
	// ...
}

func (ts *TemperatureSensor) ReadSensorValue(sensorID int) float64 {
	// Logic to read temperature from the sensor
	return 25.5
}

// PressureSensor implements SensorService for pressure measurements
type PressureSensor struct {
	// ...
}

func (ps *PressureSensor) ReadSensorValue(sensorID int) float64 {
	// Logic to read pressure from the sensor
	return 1000.2
}

// ControllerSystem handles manufacturing process control
type ControllerSystem struct {
	sensorService SensorService
}

func (cs *ControllerSystem) ProcessSensorData(sensorID int) {
	sensorValue := cs.sensorService.ReadSensorValue(sensorID)
	fmt.Printf("Sensor %d value: %.2f\n", sensorID, sensorValue)
	// Logic to analyze sensor data and control the manufacturing process
}

func main() {
	temperatureSensor := &TemperatureSensor{}
	pressureSensor := &PressureSensor{}

	controller := &ControllerSystem{sensorService: temperatureSensor}
	controller.ProcessSensorData(1)

	controller.sensorService = pressureSensor
	controller.ProcessSensorData(2)
}
