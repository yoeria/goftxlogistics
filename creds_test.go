package main

import (
	"strings"
	"testing"
)

func TestLoadCreds(t *testing.T) {
	type args struct {
		root string
	}
	tests := []struct {
		name             string
		args             args
		wantFileLocation string
	}{
		{
			name: "Print FTX key and secret",
			args: args{root: "./"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFileLocation := LoadCreds(tt.args.root); gotFileLocation != tt.wantFileLocation {
				if strings.Contains(gotFileLocation, "creds.env") {
					return
				} else {
					t.Errorf("LoadCreds() = %v, want %v", gotFileLocation, tt.wantFileLocation)
				}
			}
		})
	}
}
