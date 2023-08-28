package main

import "fmt"

// Observer interface
type Observer interface {
	Update(sensorName string, value float64)
}

// Concrete Observer: EmailAlert
type EmailAlert struct{}

func (ea *EmailAlert) Update(sensorName string, value float64) {
	if value > THRESHOLD {
		fmt.Printf("Sending email alert for %s: Value %.2f exceeded threshold\n", sensorName, value)
	}
}

// Concrete Observer: SMSAlert
type SMSAlert struct{}

func (sa *SMSAlert) Update(sensorName string, value float64) {
	if value > THRESHOLD {
		fmt.Printf("Sending SMS alert for %s: Value %.2f exceeded threshold\n", sensorName, value)
	}
}

// Concrete Observer: NotificationAlert
type NotificationAlert struct{}

func (na *NotificationAlert) Update(sensorName string, value float64) {
	if value > THRESHOLD {
		fmt.Printf("Sending notification for %s: Value %.2f exceeded threshold\n", sensorName, value)
	}
}

// Subject (Observable)
type SensorData struct {
	observers []Observer
}

func (sd *SensorData) AddObserver(observer Observer) {
	sd.observers = append(sd.observers, observer)
}

func (sd *SensorData) NotifyObservers(sensorName string, value float64) {
	for _, observer := range sd.observers {
		observer.Update(sensorName, value)
	}
}

func (sd *SensorData) SetSensorValue(sensorName string, value float64) {
	sd.NotifyObservers(sensorName, value)
}

const THRESHOLD = 100.0

func main() {
	sensorData := &SensorData{}
	emailAlert := &EmailAlert{}
	smsAlert := &SMSAlert{}
	notificationAlert := &NotificationAlert{}

	sensorData.AddObserver(emailAlert)
	sensorData.AddObserver(smsAlert)
	sensorData.AddObserver(notificationAlert)

	sensorData.SetSensorValue("Air Quality", 120.0)
	sensorData.SetSensorValue("Noise Level", 80.0)
}
