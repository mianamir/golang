package main

import (
	"fmt"
)

type DataStorageStrategy interface {
	StoreData(data map[string]float64)
}

type DatabaseStorageStrategy struct{}

func (d *DatabaseStorageStrategy) StoreData(data map[string]float64) {
	fmt.Println("Storing data in the database:", data)
}

type CloudStorageStrategy struct{}

func (c *CloudStorageStrategy) StoreData(data map[string]float64) {
	fmt.Println("Storing data in cloud storage:", data)
}

type DataCollector struct {
	storageStrategy DataStorageStrategy
}

func (d *DataCollector) SetStorageStrategy(storageStrategy DataStorageStrategy) {
	d.storageStrategy = storageStrategy
}

func (d *DataCollector) CollectAndStoreData(data map[string]float64) {
	fmt.Println("Collecting data...")
	d.storageStrategy.StoreData(data)
}

func main() {
	var databaseStrategy DataStorageStrategy
	databaseStrategy = &DatabaseStorageStrategy{}

	dataCollector := DataCollector{
		storageStrategy: databaseStrategy,
	}

	sensorData := map[string]float64{
		"Temperature": 28.5,
		"Humidity":    64.2,
		"Air Quality": 30.1,
	}

	dataCollector.CollectAndStoreData(sensorData)
}
