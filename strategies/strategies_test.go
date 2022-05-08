package strategies

import (
	"testing"

	"github.com/sdcoffey/techan"
)

func TestStrategyEMA(t *testing.T) {
	type args struct {
		series *techan.TimeSeries
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			StrategyEMA(tt.args.series)
		})
	}
}
