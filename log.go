package main

import "time"

type Log struct {
	// What triggered the log entry
	Trigger string `json:"trigger"`
	Time    time.Time `json:"time"`
	
}
