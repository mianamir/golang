package inheritence_composition

import "fmt"

// Vehicle represents a basic vehicle with make and model.
type Vehicle struct {
	Make  string
	Model string
}

// Car represents a type of vehicle and has a drive method.
type Car struct {
	Vehicle
}

// Drive prints a message that the car is driving.
func (c *Car) Drive() {
	fmt.Printf("%s %s is driving.\n", c.Make, c.Model)
}

// Engine represents an engine that can be started.
type Engine struct{}

// Start prints a message that the engine has started.
func (e *Engine) Start() {
	fmt.Println("Engine started.")
}

// CarWithEngine represents a car with an engine.
type CarWithEngine struct {
	Vehicle
	Engine Engine
}

// IoTSmartCar represents an IoT-based smart car with vehicle and sensors.
type IoTSmartCar struct {
	Vehicle Vehicle
	Sensors []Sensor
}

// StartEngine starts the engine of the IoT Smart Car.
func (ic *IoTSmartCar) StartEngine() {
	ic.Vehicle.Engine.Start()
}

// Sensor represents a generic sensor.
type Sensor interface {
	Measure() float64
}

// TemperatureSensor represents a temperature sensor.
type TemperatureSensor struct{}

// Measure simulates measuring temperature and returns a value.
func (ts *TemperatureSensor) Measure() float64 {
	return 25.0 // Simulated temperature reading
}

// ProximitySensor represents a proximity sensor.
type ProximitySensor struct{}

// Measure simulates measuring distance and returns a value.
func (ps *ProximitySensor) Measure() float64 {
	return 100.0 // Simulated distance reading
}

func main() {
	// Create instances of Car and CarWithEngine
	car := Car{Vehicle: Vehicle{Make: "Toyota", Model: "Camry"}}
	car.Drive()

	carWithEngine := CarWithEngine{
		Vehicle: Vehicle{Make: "Honda", Model: "Civic"},
		Engine:  Engine{},
	}
	carWithEngine.Engine.Start()

	// Create sensors
	temperatureSensor := TemperatureSensor{}
	proximitySensor := ProximitySensor{}

	// Create an IoT Smart Car
	smartCar := IoTSmartCar{
		Vehicle: Vehicle{Make: "Tesla", Model: "Model S"},
		Sensors: []Sensor{&temperatureSensor, &proximitySensor},
	}

	// Use the IoT Smart Car
	smartCar.StartEngine()
	fmt.Printf("Temperature: %.2f°C\n", temperatureSensor.Measure())
	fmt.Printf("Distance: %.2f cm\n", proximitySensor.Measure())
}
