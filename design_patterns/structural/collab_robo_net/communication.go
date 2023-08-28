package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Define MQTTCommunication struct
type MQTTCommunication struct {
	client mqtt.Client
}

func NewMQTTCommunication(brokerAddress string, port int) *MQTTCommunication {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerAddress, port))
	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error connecting to MQTT broker: %v\n", token.Error())
		os.Exit(1)
	}
	return &MQTTCommunication{client: client}
}

func (comm *MQTTCommunication) SendMessage(topic string, message string) {
	token := comm.client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error sending MQTT message: %v\n", token.Error())
	}
}

func (comm *MQTTCommunication) ReceiveMessage(topic string, callback func(string)) {
	token := comm.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		callback(string(msg.Payload()))
	})
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("Error subscribing to MQTT topic: %v\n", token.Error())
	}
}

// Define MQTTCommunicationProxy struct
type MQTTCommunicationProxy struct {
	mqttCommunication *MQTTCommunication
}

func NewMQTTCommunicationProxy(brokerAddress string, port int) *MQTTCommunicationProxy {
	mqttComm := NewMQTTCommunication(brokerAddress, port)
	return &MQTTCommunicationProxy{mqttCommunication: mqttComm}
}

func (proxy *MQTTCommunicationProxy) SendMessage(topic string, message string) {
	// Add any additional logic before sending message
	proxy.mqttCommunication.SendMessage(topic, message)
}

func (proxy *MQTTCommunicationProxy) ReceiveMessage(topic string, callback func(string)) {
	// Add any additional logic before receiving message
	proxy.mqttCommunication.ReceiveMessage(topic, callback)
}

func main() {
	// Example usage
	brokerAddress := "localhost"
	port := 1883

	proxy := NewMQTTCommunicationProxy(brokerAddress, port)

	// Simulate sending a message
	proxy.SendMessage("test_topic", "Test message")

	// Simulate receiving a message
	proxy.ReceiveMessage("test_topic", func(message string) {
		fmt.Printf("Received message: %s\n", message)
	})
	time.Sleep(time.Second) // Allow time for receiving message
}
