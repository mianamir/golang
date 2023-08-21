package grasp_patterns

import (
	"fmt"
)

// Server represents a virtual server instance in a cloud environment
type Server struct {
	ID       int
	Name     string
	IP       string
	Status   string
	CPUUsage float64
	Memory   float64
}

// CloudServices provides operations for managing cloud resources
type CloudServices struct {
	servers     []Server
	serverCount int
}

func (cs *CloudServices) CreateServer(name, ip string) Server {
	server := Server{
		ID:       cs.serverCount,
		Name:     name,
		IP:       ip,
		Status:   "Running",
		CPUUsage: 0.0,
		Memory:   0.0,
	}
	cs.servers = append(cs.servers, server)
	cs.serverCount++
	return server
}

func (cs *CloudServices) MonitorResourceUsage(server Server, cpuUsage, memory float64) {
	server.CPUUsage = cpuUsage
	server.Memory = memory
	fmt.Printf("Server %s - CPU Usage: %.2f%%, Memory: %.2f GB\n", server.Name, cpuUsage, memory)
}

func main() {
	cloudServices := CloudServices{}

	server1 := cloudServices.CreateServer("WebServer1", "192.168.1.101")
	server2 := cloudServices.CreateServer("DatabaseServer1", "192.168.1.102")

	cloudServices.MonitorResourceUsage(server1, 20.5, 8.0)
	cloudServices.MonitorResourceUsage(server2, 40.2, 16.0)
}
