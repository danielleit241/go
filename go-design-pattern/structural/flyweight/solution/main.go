package problem

import "fmt"

type ChatMessage struct {
	Content string
	Sender  *Sender
}

type Sender struct {
	Name   string
	Avatar []byte // Some big thing in memory
}

type SenderFactory struct {
	cacheSender map[string]*Sender
}

func CreateSenderFactory() *SenderFactory {
	return &SenderFactory{cacheSender: make(map[string]*Sender)}
}

func (f *SenderFactory) GetSender(name string) *Sender {
	if sender, ok := f.cacheSender[name]; ok {
		return sender
	}
	// In a real application, we would load the avatar from disk or a database
	avatar := make([]byte, 1024*300) //300KB
	sender := &Sender{Name: name, Avatar: avatar}
	f.cacheSender[name] = sender
	return sender
}

func main() {
	factory := CreateSenderFactory()
	messages := []ChatMessage{
		{Content: "hi", Sender: factory.GetSender("Peter")},
		{Content: "oh, here you are", Sender: factory.GetSender("Mary")},
		{Content: "how are you doing?", Sender: factory.GetSender("Peter")},
		{Content: "fine, thanks", Sender: factory.GetSender("Mary")},
	}
	fmt.Println(messages)

	// We just created 4 messages, but we have 2 senders, so we have 2 big avatars in memory that are shared.
	// Total memory used: 300KB + 400KB = 700KB
}
