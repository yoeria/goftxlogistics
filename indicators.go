package main

import (
	talib "github.com/markcheno/go-talib"
	"github.com/sdcoffey/techan"
)

// Format for delivering the required data for calculating the indicators.
type StratData struct {
	techanSeries *techan.TimeSeries
	inReal []float64
	timeFrame int
	kPeriod int
	dPeriod int
	maType talib.MaType
}

func (s StratData) calcStochRSI() ([]float64,[]float64){
	return talib.StochRsi(s.inReal, s.timeFrame, s.kPeriod, s.dPeriod, s.maType)
}

// Calculate FAST Stochastic values
func (s StratData) calcFastStochastic() techan.Indicator{
	return techan.NewFastStochasticIndicator(s.techanSeries, s.timeFrame)
}
