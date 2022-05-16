package main

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestLoadCreds(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "print FTX key and secret",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			wg.Add(2)
			LoadCreds("./creds.env")
			go fmt.Print(ReadOnlyKey, ReadOnlySecret)
			wg.Wait()
		})
	}
}
