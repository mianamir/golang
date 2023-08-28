package singleton

import "sync"

type FactoryDashboard struct {
	sensors []string
}

var instance *FactoryDashboard
var once sync.Once

func GetFactoryDashboard() *FactoryDashboard {
	once.Do(func() {
		instance = &FactoryDashboard{}
	})
	return instance
}

func (d *FactoryDashboard) AddSensor(sensor string) {
	d.sensors = append(d.sensors, sensor)
}

func (d *FactoryDashboard) GetSensors() []string {
	return d.sensors
}
