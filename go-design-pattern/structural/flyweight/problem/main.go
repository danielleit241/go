package problem

import "fmt"

type ChatMessage struct {
	Content      string
	SenderName   string
	SenderAvatar []byte // Some big thing in memory
}

func main() {
	fmt.Println([]ChatMessage{
		{
			Content:      "hi",
			SenderName:   "Peter",
			SenderAvatar: make([]byte, 1024*300), //300KB
		},
		{
			Content:      "oh, here you are",
			SenderName:   "Mary",
			SenderAvatar: make([]byte, 1024*400), //400KB
		},
		{
			Content:      "how are you doing?",
			SenderName:   "Peter",
			SenderAvatar: make([]byte, 1024*300), //300KB
		},
		{
			Content:      "fine, thanks",
			SenderName:   "Mary",
			SenderAvatar: make([]byte, 1024*400), //400KB
		},
	})

	// We just created 4 messages, but we have 2 senders, so we have 2 big avatars in memory that are duplicated.
	// Total memory used: 300KB + 400KB + 300KB + 400KB = 1.4MB
}
