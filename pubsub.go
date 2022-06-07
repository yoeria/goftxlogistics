// This file is meant for creating channels that send and receive messages from the redis server by subscribing. When a value on the db is changed, the client will be notified by one of these channels
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Use this subscription func to subscribe to messages related to the "bot" preferences + other channels
func SubToChannel(channel string) (ch <-chan *redis.Message) {
	pubsub := rdb.Subscribe(ctx, channel)

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch = pubsub.Channel()

	// When pubsub is closed channel is closed too.
	time.AfterFunc(time.Second, func() {

		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Printf("Time:\t%v\nrdbChannel:\t%v\nMessage:\t%v\n",time.Now().Local().Format(time.RFC822), msg.Channel, msg.Payload)
	}
	return
}

// Publish a message to the rdb pubsub channel given as an argument. Message as the second argument.
func PubToChannel(channel string, message interface{}) {
	// In configuration interface:
	err = rdb.Publish(ctx, channel, message).Err()
	if err != nil {
		// Print the error to stdout
		fmt.Printf("%v", err.Error())
		panic(err)
	}
}
