package problem

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Send message: %s via email\n", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) {
	fmt.Printf("Send message: %s via SMS\n", message)
}

// If we want to send notifications via both email and SMS, we would need to create a new struct that combines both notifiers

type EmailSMSNotifier struct {
	emailNotifier *EmailNotifier
	smsNotifier   *SMSNotifier
}

func (e *EmailSMSNotifier) Send(message string) {
	e.emailNotifier.Send(message)
	e.smsNotifier.Send(message)
}

type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.Send(message)
}

func main() {
	emailSMSNotifier := &EmailSMSNotifier{
		emailNotifier: &EmailNotifier{},
		smsNotifier:   &SMSNotifier{},
	}

	s := NotificationService{
		notifier: emailSMSNotifier,
	}

	s.SendNotification("Hello, World!")
}
