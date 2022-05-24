package main

import (
	"context"
	"time"
)

type Log struct {
	// What triggered the log entry
	Trigger string `json:"trigger"`
	// Time of logging
	Time time.Time `json:"time"`
	// Request message sent to server [proto]
	Request []byte `json:"request"`
	// Response received back from server [proto]
	Response []byte `json:"response"`
}

func (l *Log) Send() {
	rdb.MSetNX(context.Background(), l)
}
