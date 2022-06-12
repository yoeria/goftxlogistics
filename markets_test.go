package main

import (
	"reflect"
	"testing"
)

func TestCurrentPotentFutures(t *testing.T) {
	type args struct {
		perpetual bool
	}
	tests := []struct {
		name     string
		args     args
		wantList []*TradableFutures
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotList := CurrentPotentFutures(tt.args.perpetual); !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("CurrentPotentFutures() = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
