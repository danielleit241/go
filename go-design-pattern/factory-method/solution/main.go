package solution

import "fmt"

// Implementing the Strategy pattern to adhere to the Open/Closed Principle.
type Notifier interface {
	Send(message string)
}

// Concrete strategy for email notification
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Send message: %s via email\n", message)
}

// Concrete strategy for SMS notification
type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) {
	fmt.Printf("Send message: %s via SMS\n", message)
}

// We can easily add new notification types without modifying existing code by implementing the Notifier interface.
type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.Send(message)
}

// Factory method to create the notifier based on the type of notification.
func CreateNotifier(notificationType string) Notifier {
	switch notificationType {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SMSNotifier{}
	default:
		return nil
	}
}

func main() {
	s := NotificationService{
		notifier: CreateNotifier("email"),
	}

	s.SendNotification("Hello, World!")
}
