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
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{}

	notificationService := &NotificationService{notifier: emailNotifier}
	notificationService.SendNotification("Hello, World!")

	notificationService.notifier = smsNotifier
	notificationService.SendNotification("Hello, World!")
}

//Pros:
//1. Adheres to the Open/Closed Principle: We can add new notification types without modifying existing code.
//2. Promotes code reusability and separation of concerns by encapsulating the notification logic in separate classes.
//3. Makes the code more flexible and easier to maintain.

//Cons:
//1. Introduces additional complexity by requiring the creation of multiple classes for each notification type.
//2. May require more memory and resources due to the use of interfaces and multiple struct types.
//3. Can lead to a proliferation of classes if there are many notification types, which may make the codebase harder to navigate.
