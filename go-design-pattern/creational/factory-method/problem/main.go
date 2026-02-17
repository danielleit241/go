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

func main() {
	s := NotificationService{
		// I don't want my user init the notifier, I want to hide the implementation details of the notifier from the user.
		// They should call something to get the notifier, and then use it to send the notification.
		// CreateNotifier is a factory method that creates the notifier based on the type of notification.
		notifier: &EmailNotifier{},
	}

	s.SendNotification("Hello, World!")
}
