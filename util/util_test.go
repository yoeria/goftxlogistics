package util

import (
	"testing"
)

func TestGetTradeAmount(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTradeAmount(); got != tt.want {
				t.Errorf("GetTradeAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
