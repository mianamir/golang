package main

import "fmt"

// Adapter Pattern: Actuator
// This pattern helps integrate different actuator types into a common interface.
type Actuator interface {
	PerformAction(action string) string
}

type RoboticArm struct{}

func (ra *RoboticArm) PerformAction(action string) string {
	return fmt.Sprintf("Robotic arm performing action: %s", action)
}

// Decorator Pattern: Actuator
// This pattern enhances the functionality of actuators by adding extra behaviors.
type ActuatorDecorator struct {
	actuator Actuator
}

func (ad *ActuatorDecorator) PerformAction(action string) string {
	return ""
}

type EnhancedRoboticArm struct {
	ActuatorDecorator
}

func (era *EnhancedRoboticArm) PerformAction(action string) string {
	baseActionResult := era.actuator.PerformAction(action)
	enhancedActionResult := fmt.Sprintf("Enhanced %s", baseActionResult)
	return enhancedActionResult
}

// Adapter Pattern: Sensor
// This pattern helps integrate different sensor types into a common interface.
type Sensor interface {
	Measure() string
}

type CameraSensor struct{}

func (cs *CameraSensor) Measure() string {
	return "Image data from camera sensor"
}

type ProximitySensor struct{}

func (ps *ProximitySensor) Measure() string {
	return "Proximity data from proximity sensor"
}

// Decorator Pattern: Sensor
// This pattern enhances the functionality of sensors by adding extra behaviors.
type SensorDecorator struct {
	sensor Sensor
}

func (sd *SensorDecorator) Measure() string {
	return ""
}

type EnhancedCameraSensor struct {
	SensorDecorator
}

func (ecs *EnhancedCameraSensor) Measure() string {
	baseMeasurement := ecs.sensor.Measure()
	enhancedMeasurement := fmt.Sprintf("Enhanced %s", baseMeasurement)
	return enhancedMeasurement
}

type EnhancedProximitySensor struct {
	SensorDecorator
}

func (eps *EnhancedProximitySensor) Measure() string {
	baseMeasurement := eps.sensor.Measure()
	enhancedMeasurement := fmt.Sprintf("Enhanced %s", baseMeasurement)
	return enhancedMeasurement
}

func main() {
	// Example usage
	roboticArm := &RoboticArm{}
	enhancedRoboticArm := &EnhancedRoboticArm{ActuatorDecorator: ActuatorDecorator{actuator: roboticArm}}

	cameraSensor := &CameraSensor{}
	enhancedCameraSensor := &EnhancedCameraSensor{SensorDecorator: SensorDecorator{sensor: cameraSensor}}

	proximitySensor := &ProximitySensor{}
	enhancedProximitySensor := &EnhancedProximitySensor{SensorDecorator: SensorDecorator{sensor: proximitySensor}}

	actionResult := enhancedRoboticArm.PerformAction("Pick and place")
	measurement := enhancedCameraSensor.Measure()
	proximityMeasurement := enhancedProximitySensor.Measure()

	fmt.Println(actionResult)
	fmt.Println(measurement)
	fmt.Println(proximityMeasurement)
}
