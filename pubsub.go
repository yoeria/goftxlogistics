// This file is meant for creating channels that send and receive messages from the redis server by subscribing. When a value on the db is changed, the client will be notified by one of these channels
package main

import (
	"fmt"
	"time"
)

// Use this subscription to subscribe to messages related to the "bot" preferences
func pubsub() {

	pubsub := rdb.Subscribe(ctx, "mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	err = rdb.Publish(ctx, "mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
