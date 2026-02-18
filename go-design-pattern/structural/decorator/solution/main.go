package main

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

type TelegramNotifier struct{}

func (t *TelegramNotifier) Send(message string) {
	fmt.Printf("Send message: %s via Telegram\n", message)
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

func (d *NotifierDecorator) Send(message string) {
	d.notifier.Send(message)

	if d.core != nil {
		d.core.Send(message)
	}
}

// Like add to stack

func (nd NotifierDecorator) Decorate(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		notifier: notifier,
	}
}

type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(&EmailNotifier{}).
		Decorate(&SMSNotifier{}).
		Decorate(&TelegramNotifier{})

	s := NotificationService{
		notifier: &notifier,
	}

	s.SendNotification("Hello, World!")
}
