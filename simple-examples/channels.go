package main

import "fmt"

// main demonstrates a simple channel
// Note: by default, sends and receives block until both the sender
// and the receiver are ready.
func main() {
	messages := make(chan string)

	// Send message to the channel from a new goroutine
	go func() { messages <- "hello" }()

	msg := <-messages
	fmt.Println(msg)
}
