package main

import (
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

// Converts received candle data from exchange to format expected in techan lib
func ParseSeries(data []markets.Candle) (series *techan.TimeSeries) {
	// Shorter declaration
	sc := series.Candles
	d := data
	// Move values to time series
	for i := range d {
		sc[i].OpenPrice = big.NewDecimal(d[i].Open)
		sc[i].ClosePrice = big.NewDecimal(d[i].Close)
		sc[i].MinPrice = big.NewDecimal(d[i].Low)
		sc[i].MaxPrice = big.NewDecimal(d[i].High)
		sc[i].Volume = big.NewDecimal(d[i].Volume)
		sc[i].TradeCount = uint(GetTradeAmount())
	}
	return
}

// Gets the close price of each candle and assign that as data value in index
func ParseInReal(data []markets.Candle) (inReal []float64) {
	for i := range data {
		inReal[i] = data[i].Close
	}
	return
}
