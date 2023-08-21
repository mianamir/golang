package grasp_patterns

import (
	"fmt"
)

// NotificationService defines the interface for sending notifications
type NotificationService interface {
	SendNotification(message string)
}

// EmailNotification implements NotificationService
type EmailNotification struct {
	// ...
}

func (en *EmailNotification) SendNotification(message string) {
	// Logic to send an email notification
	fmt.Printf("Email notification sent: %s\n", message)
}

// SMSNotification implements NotificationService
type SMSNotification struct {
	// ...
}

func (sn *SMSNotification) SendNotification(message string) {
	// Logic to send an SMS notification
	fmt.Printf("SMS notification sent: %s\n", message)
}

func main() {
	notificationService := &EmailNotification{}
	sendNotification(notificationService, "Your order has been shipped!")

	notificationService = &SMSNotification{}
	sendNotification(notificationService, "Your payment is due soon!")
}

func sendNotification(service NotificationService, message string) {
	service.SendNotification(message)
}
