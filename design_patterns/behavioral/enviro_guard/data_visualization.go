package main

import (
	"fmt"
)

type Observer interface {
	Update(sensorName string, value float64)
}

type Sensor struct {
	Name      string
	Value     float64
	Observers []Observer
}

func NewSensor(name string) *Sensor {
	return &Sensor{
		Name:      name,
		Value:     0,
		Observers: make([]Observer, 0),
	}
}

func (s *Sensor) AddObserver(observer Observer) {
	s.Observers = append(s.Observers, observer)
}

func (s *Sensor) RemoveObserver(observer Observer) {
	index := -1
	for i, obs := range s.Observers {
		if obs == observer {
			index = i
			break
		}
	}

	if index >= 0 {
		s.Observers = append(s.Observers[:index], s.Observers[index+1:]...)
	}
}

func (s *Sensor) SetValue(value float64) {
	s.Value = value
	s.NotifyObservers()
}

func (s *Sensor) NotifyObservers() {
	for _, observer := range s.Observers {
		observer.Update(s.Name, s.Value)
	}
}

type DataVisualizer struct{}

func (dv *DataVisualizer) Update(sensorName string, value float64) {
	fmt.Printf("Visualizing %s value: %f\n", sensorName, value)
}

func main() {
	temperatureSensor := NewSensor("Temperature Sensor")
	airQualitySensor := NewSensor("Air Quality Sensor")

	dataVisualizer := &DataVisualizer{}

	temperatureSensor.AddObserver(dataVisualizer)
	airQualitySensor.AddObserver(dataVisualizer)

	// Simulate data updates and visualization
	temperatureSensor.SetValue(25.7)
	airQualitySensor.SetValue(38.2)
}
