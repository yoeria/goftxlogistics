package main

import (
	"fmt"
	"testing"
)

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
			LoadCreds("./")
			go fmt.Print(ReadOnlyKey, ReadOnlySecret)
			wg.Wait()
		})
	}
}
