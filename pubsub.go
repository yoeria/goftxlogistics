// This file is meant for creating channels that send and receive messages from the redis server by subscribing. When a value on the db is changed, the client will be notified by one of these channels
package main

import (
	"fmt"
	"time"
)


// Use this subscription to subscribe to messages related to the "bot" preferences
	// Wait for confirmation that subscription is created before publishing anything.
	// Go channel which receives messages.
	// Publish a message.
	// When pubsub is closed channel is closed too.
	// Consume messages.
func SubToPreferences() {
	pubsub := rdb.Subscribe(ctx, "preferences")

	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	// In preferences interface:
	err = rdb.Publish(ctx, "preferences", "hello").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {

		_ = pubsub.Close()
	})

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
