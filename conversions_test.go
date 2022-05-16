package main

import (
	"reflect"
	"testing"

	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/sdcoffey/techan"
)

func TestParseSeries(t *testing.T) {
	type args struct {
		data []markets.Candle
	}
	tests := []struct {
		name       string
		args       args
		wantSeries *techan.TimeSeries
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSeries := ParseSeries(tt.args.data); !reflect.DeepEqual(gotSeries, tt.wantSeries) {
				t.Errorf("ParseSeries() = %v, want %v", gotSeries, tt.wantSeries)
			}
		})
	}
}
