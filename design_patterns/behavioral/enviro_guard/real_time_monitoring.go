package main

import "fmt"

// Observer defines the interface for observers.
type Observer interface {
	Update(sensorName string, value interface{})
}

// Sensor represents a subject that produces values.
type Sensor struct {
	name      string
	value     interface{}
	observers []Observer
}

func NewSensor(name string) *Sensor {
	return &Sensor{
		name:      name,
		value:     nil,
		observers: make([]Observer, 0),
	}
}

func (s *Sensor) AddObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Sensor) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Sensor) SetValue(value interface{}) {
	s.value = value
	s.NotifyObservers()
}

func (s *Sensor) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.name, s.value)
	}
}

// RealTimeMonitor represents a concrete observer class for real-time monitoring.
type RealTimeMonitor struct{}

func (rtm *RealTimeMonitor) Update(sensorName string, value interface{}) {
	fmt.Printf("Real-time monitoring: %s value: %v\n", sensorName, value)
}

func main() {
	temperatureSensor := NewSensor("Temperature Sensor")
	humiditySensor := NewSensor("Humidity Sensor")

	realTimeMonitor := &RealTimeMonitor{}

	temperatureSensor.AddObserver(realTimeMonitor)
	humiditySensor.AddObserver(realTimeMonitor)

	// Simulate data updates and real-time monitoring
	temperatureSensor.SetValue(26.8)
	humiditySensor.SetValue(57.3)
}
