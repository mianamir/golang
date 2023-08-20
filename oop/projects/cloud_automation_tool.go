package main

import "fmt"

// CloudService is the interface defining the contract for provisioning a cloud service.
type CloudService interface {
	Provision()
}

// VirtualMachine represents a virtual machine cloud service.
type VirtualMachine struct {
	name   string
	cpu    int
	memory int
}

// Provision implements the CloudService interface for virtual machine provisioning.
func (vm *VirtualMachine) Provision() {
	fmt.Printf("Provisioned Virtual Machine: %s, CPU: %d, Memory: %d GB\n", vm.name, vm.cpu, vm.memory)
}

// DatabaseInstance represents a database instance cloud service.
type DatabaseInstance struct {
	name    string
	storage int
	engine  string
}

// Provision implements the CloudService interface for database instance provisioning.
func (db *DatabaseInstance) Provision() {
	fmt.Printf("Provisioned Database Instance: %s, Storage: %d GB, Engine: %s\n", db.name, db.storage, db.engine)
}

// DeploymentStrategy is the interface defining the contract for deploying a cloud service.
type DeploymentStrategy interface {
	Deploy()
}

// BlueGreenDeployment represents a blue-green deployment strategy.
type BlueGreenDeployment struct{}

// Deploy implements the DeploymentStrategy interface for blue-green deployment.
func (bg *BlueGreenDeployment) Deploy() {
	fmt.Println("Performing Blue-Green Deployment")
}

// CanaryDeployment represents a canary deployment strategy.
type CanaryDeployment struct{}

// Deploy implements the DeploymentStrategy interface for canary deployment.
func (cd *CanaryDeployment) Deploy() {
	fmt.Println("Performing Canary Deployment")
}

// CloudAutomationSystem represents the facade for automating cloud service provisioning and deployment.
type CloudAutomationSystem struct {
	cloudService       CloudService
	deploymentStrategy DeploymentStrategy
}

// Automate automates cloud service provisioning and deployment.
func (cas *CloudAutomationSystem) Automate() {
	cas.cloudService.Provision()
	cas.deploymentStrategy.Deploy()
}

// MonitoringObserver is the interface defining the contract for receiving monitoring alerts.
type MonitoringObserver interface {
	NotifyAlert(alert string)
}

// AlertNotifier notifies alerts via a monitoring observer.
type AlertNotifier struct {
	observer MonitoringObserver
}

// NotifyAlert notifies an alert via the monitoring observer.
func (an *AlertNotifier) NotifyAlert(alert string) {
	an.observer.NotifyAlert(alert)
}

// EmailAlertObserver represents an email-based monitoring alert observer.
type EmailAlertObserver struct{}

// NotifyAlert sends an email notification for a monitoring alert.
func (ea *EmailAlertObserver) NotifyAlert(alert string) {
	fmt.Printf("Alert: %s - Notified via Email\n", alert)
}

func main() {
	// Creating cloud services
	vm := &VirtualMachine{name: "WebServer", cpu: 2, memory: 8}
	db := &DatabaseInstance{name: "MySQLDB", storage: 100, engine: "MySQL"}

	// Creating deployment strategies
	blueGreenDeployment := &BlueGreenDeployment{}
	canaryDeployment := &CanaryDeployment{}

	// Creating cloud automation system
	cloudAutomation := &CloudAutomationSystem{cloudService: vm, deploymentStrategy: blueGreenDeployment}

	// Creating alert notifier and observer
	alertNotifier := &AlertNotifier{observer: &EmailAlertObserver{}}

	// Automating cloud deployment
	cloudAutomation.Automate()

	// Notifying alerts
	alertNotifier.NotifyAlert("High CPU Usage")
}
