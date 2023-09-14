package main

import (
	"fmt"
	"reflect"
)

type Connector interface {
    Connect()
}

type DatabaseConnector struct{}

func (db DatabaseConnector) Connect() {
    fmt.Println("Connected to the database.")
}

type APIConnector struct{}

func (api APIConnector) Connect() {
    fmt.Println("Connected to the API.")
}

func createConnector(userInput string) (Connector, error) {
    switch userInput {
    case "database":
        return DatabaseConnector{}, nil
    case "api":
        return APIConnector{}, nil
    default:
        return nil, fmt.Errorf("Invalid data source")
    }
}

type Machine struct {
    MachineID string
    Status    bool
}

func (m *Machine) Start() {
    m.Status = true
    fmt.Printf("Machine %s started.\n", m.MachineID)
}

func (m *Machine) Stop() {
    m.Status = false
    fmt.Printf("Machine %s stopped.\n", m.MachineID)
}

func (m *Machine) StatusInfo() string {
    status := "stopped"
    if m.Status {
        status = "running"
    }
    return fmt.Sprintf("Machine %s is %s.", m.MachineID, status)
}

type FactoryController struct {
    Machines map[string]*Machine
}

func NewFactoryController() *FactoryController {
    return &FactoryController{
        Machines: make(map[string]*Machine),
    }
}

func (f *FactoryController) AddMachine(machineID string) {
    machine := &Machine{
        MachineID: machineID,
        Status:    false,
    }
    f.Machines[machineID] = machine
    fmt.Printf("Machine %s added to the factory.\n", machineID)
}

func (f *FactoryController) OperateMachine(machineID string, action string) {
    machine, exists := f.Machines[machineID]
    if !exists {
        fmt.Println("Machine not found.")
        return
    }

    // Using reflection to dynamically call methods based on user action
    method := reflect.ValueOf(machine).MethodByName(action)
    if !method.IsValid() {
        fmt.Println("Invalid action. Use 'Start' or 'Stop'.")
        return
    }

    method.Call([]reflect.Value{})
}

func (f *FactoryController) ListMachines() {
    fmt.Println("Current machine status:")
    for machineID, machine := range f.Machines {
        fmt.Println(machine.StatusInfo())
    }
}

func main() {
    userInput := "database" // This could also be "api" based on user choice

    connector, err := createConnector(userInput)
    if err != nil {
        fmt.Println(err)
        return
    }

    connector.Connect()

    factory := NewFactoryController()

    factory.AddMachine("Machine-1")
    factory.AddMachine("Machine-2")

    factory.OperateMachine("Machine-1", "Start")
    factory.OperateMachine("Machine-2", "Stop")

    factory.ListMachines()
}
