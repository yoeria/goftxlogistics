package main

import (
	"testing"

	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/sdcoffey/techan"
)

func TestStrategyEMA(t *testing.T) {
	candles,err:= rc.Candles(&markets.RequestForCandles{ProductCode: "BTC-PERP",Resolution: (5*60)})
	if err != nil {
		panic(err)
	}

	tradeRecord:= techan.NewTradingRecord()

	type args struct {
		series *techan.TimeSeries
		record *techan.TradingRecord
		Signal chan Order
	}
	tests := []struct {
		name string
		args args
	}{{name: "TestStrategyEMA", args: args{series: ParseSeries(*candles), record: tradeRecord, Signal: make(chan Order)}}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StrategyEMA(tt.args.series, tt.args.record, tt.args.Signal)
			t.Log(tt.args.record.Trades)
		})
	}
}
