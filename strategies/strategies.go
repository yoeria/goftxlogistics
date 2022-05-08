package strategies

import (
	talib "github.com/markcheno/go-talib"
	t "github.com/sdcoffey/techan"
)

// Indicator average ratio of at least 0.8 (80%) should start buying cycle
// If calculateSuccessRatio >= 0.8 -> trigger an order
func calculateSuccessRatio(_ chan int) {

	for {
		select {}
	}

}

// Calculate FAST Stochastic values
func calcFastStoch(series *t.TimeSeries, timeframe int) {
	t.NewFastStochasticIndicator(series, timeframe)

}

func calcStochRSI(inReal []float64, timeframe, kPeriod, dPeriod int, maType talib.MaType) ([]float64, []float64) {
	return talib.StochRsi(inReal, timeframe, kPeriod, dPeriod, maType)
}

/*
5-8-13 EMA Daytrading strategy
 https://www.investopedia.com/articles/active-trading/010116/perfect-moving-averages-day-trading.asp
*/
func StrategyEMA(series *t.TimeSeries) {
	closePriceIndicator := t.NewClosePriceIndicator(series)
	indicator5 := t.NewEMAIndicator(closePriceIndicator, 5)
//	indicator8 := t.NewEMAIndicator(closePriceIndicator, 8)
//	indicator13 := t.NewEMAIndicator(closePriceIndicator, 13)

	// record trades on this object
	record := t.NewTradingRecord()

	entryConstant := t.NewConstantIndicator(30)
	exitConstant := t.NewConstantIndicator(10)

	// Is satisfied when the price ema moves above 30 and the current position is new
	entryRule := t.And(
		t.NewCrossUpIndicatorRule(entryConstant, indicator5),
		t.PositionNewRule{})

	// Is satisfied when the price ema moves below 10 and the current position is open
	exitRule := t.And(
		t.NewCrossDownIndicatorRule(indicator5, exitConstant),
		t.PositionOpenRule{})

	strategy := t.RuleStrategy{
		UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	strategy.ShouldEnter(0, record) // returns
	strategy.ShouldExit(0,record)
}
