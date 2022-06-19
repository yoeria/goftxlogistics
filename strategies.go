package main

import (
	t "github.com/sdcoffey/techan"
)

// Indicator average ratio of at least 0.8 (80%) should start buying cycle
// If calculateSuccessRatio >= 0.8 -> trigger an order
func calculateSuccessRatio(_ chan int) {

	for {
		select {}
	}

}

/*
5-8-13 EMA Daytrading strategy
 https://www.investopedia.com/articles/active-trading/010116/perfect-moving-averages-day-trading.asp
*/
func StrategyEMA(series *t.TimeSeries, record *t.TradingRecord) {
	closePriceIndicator := t.NewClosePriceIndicator(series)
	indicator5 := t.NewEMAIndicator(closePriceIndicator, 5)
	//	indicator8 := t.NewEMAIndicator(closePriceIndicator, 8)
	//	indicator13 := t.NewEMAIndicator(closePriceIndicator, 13)

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

	entryTrigger := strategy.ShouldEnter(0, record) // returns
	exitTrigger := strategy.ShouldExit(0, record)   // returns

	if entryTrigger {
	}

	if exitTrigger {

	}

}

func templateStrategy(series *t.TimeSeries) {
	indicator := t.NewClosePriceIndicator(series)

	// record trades on this object
	record := t.NewTradingRecord()

	entryConstant := t.NewConstantIndicator(30)
	exitConstant := t.NewConstantIndicator(10)

	// Is satisfied when the price ema moves above 30 and the current position is new
	entryRule := t.And(
		t.NewCrossUpIndicatorRule(entryConstant, indicator),
		t.PositionNewRule{})

	// Is satisfied when the price ema moves below 10 and the current position is open
	exitRule := t.And(
		t.NewCrossDownIndicatorRule(indicator, exitConstant),
		t.PositionOpenRule{})

	strategy := t.RuleStrategy{
		UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	strategy.ShouldEnter(0, record) // returns
}
