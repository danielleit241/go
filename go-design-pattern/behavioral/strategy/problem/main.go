package problem

import "fmt"

type Notification struct {
	NotificationType string
}

func (n *Notification) Send(message string) {
	switch n.NotificationType {
	case "email":
		fmt.Printf("Send message: %s via email\n", message)
	case "sms":
		fmt.Printf("Send message: %s via SMS\n", message)
	}

	// If we want to add a new notification type, we need to modify the Send method, which violates the Open/Closed Principle.
}

func main() {
	s := &Notification{NotificationType: "email"}
	s.Send("Hello, World!")
}
