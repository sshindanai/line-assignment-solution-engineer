package app

import "github.com/r3labs/sse"

func Publish() {
	server.Publish("messages", &sse.Event{
		Data: []byte("ping"),
	})
}
