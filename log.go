package main

import (
	"context"
	"time"
)

type Log struct {
	// What triggered the log entry
	Trigger string    `json:"trigger"`
	Time    time.Time `json:"time"`
}

func (l *Log) Send() {
	RDB.MSetNX(context.Background(), l)
}
