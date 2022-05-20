package main

import (
	"testing"
	"time"
)

func TestLog_Send(t *testing.T) {
	type fields struct {
		Trigger string
		Time    time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
struct{name string; fields fields}{name: "Send Log struct as set",fields: fields{Trigger: STRATEGY,Time: time.Now()}}	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Trigger: tt.fields.Trigger,
				Time:    tt.fields.Time,
			}
			l.Send()
		})
	}
}
