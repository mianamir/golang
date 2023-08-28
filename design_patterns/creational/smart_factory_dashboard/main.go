package main

import (
    "../design-patterns-golang/abstractfactory"
    "../design-patterns-golang/builder"
    "../design-patterns-golang/factorymethod"
    "../design-patterns-golang/objectpool"
    "../design-patterns-golang/prototype"
    "../design-patterns-golang/singleton"
    "fmt"
)

func main() {
	// Factory Method Pattern
	tempSensorFactory := &factorymethod.TemperatureSensorFactory{}
	tempSensorController := NewSensorController(tempSensorFactory)
	fmt.Println(tempSensorController.GetMeasurement())

	// Abstract Factory Pattern
	cuttingMachineFactory := &abstractfactory.CuttingMachineFactory{}
	cuttingMachineDashboard := NewDashboard(cuttingMachineFactory)
	fmt.Println(cuttingMachineDashboard.ShowMachine())

	// Singleton Pattern
	dashboard1 := singleton.GetFactoryDashboard()
	dashboard1.AddSensor("Temperature Sensor")
	dashboard1.AddSensor("Pressure Sensor")

	dashboard2 := singleton.GetFactoryDashboard()
	dashboard2.AddSensor("Humidity Sensor")

	fmt.Println(dashboard1.GetSensors())
	fmt.Println(dashboard2.GetSensors())

	// Builder Pattern
	machineBuilder := builder.NewComplexMachineBuilder()
	machineBuilder.AddName("Robotic Assembly Line")
	machineBuilder.AddSensors([]string{"Temperature Sensor", "Pressure Sensor", "Humidity Sensor"})

	complexMachine := machineBuilder.GetMachine()
	fmt.Println(complexMachine.Display())

	// Prototype Pattern
	basicMachine := &prototype.BasicMachine{}
	basicMachine.SetSensors([]string{"Temperature Sensor"})

	clonedMachine := basicMachine.Clone().(*builder.MachineImpl)
	clonedMachine.SetName("Cloned Machine")
	clonedMachine.SetSensors([]string{"Pressure Sensor"})

	fmt.Println(basicMachine.Display())
	fmt.Println(clonedMachine.Display())

	// Object Pool Pattern
	machinePool := objectpool.NewMachinePool(3)

	machine1 := machinePool.GetMachine().(*builder.MachineImpl)
	machine1.SetName("Machine 1")
	machine1.SetSensors([]string{"Temperature Sensor"})

	machine2 := machinePool.GetMachine().(*builder.MachineImpl)
	machine2.SetName("Machine 2")
	machine2.SetSensors([]string{"Pressure Sensor"})

	machine3 := machinePool.GetMachine().(*builder.MachineImpl)
	machine3.SetName("Machine 3")
	machine3.SetSensors([]string{"Humidity Sensor"})

	machine4 := machinePool.GetMachine().(*builder.MachineImpl)
	machine4.SetName("Machine 4")
	machine4.SetSensors([]string{"Vibration Sensor"})

	fmt.Println(machine1.Display())
	fmt.Println(machine2.Display())
	fmt.Println(machine3.Display())
	fmt.Println(machine4.Display())
}
