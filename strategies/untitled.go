package strategies

import "github.com/sdcoffey/techan"

func newStrategy(series *techan.TimeSeries){
indicator := techan.NewClosePriceIndicator(series)

// record trades on this object
record := techan.NewTradingRecord()

entryConstant := techan.NewConstantIndicator(30)
exitConstant := techan.NewConstantIndicator(10)

// Is satisfied when the price ema moves above 30 and the current position is new
entryRule := techan.And(
	techan.NewCrossUpIndicatorRule(entryConstant, indicator),
		techan.PositionNewRule{})

			// Is satisfied when the price ema moves below 10 and the current position is open
			exitRule := techan.And(
				techan.NewCrossDownIndicatorRule(indicator, exitConstant),
					techan.PositionOpenRule{})

					strategy := techan.RuleStrategy{
						UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
							EntryRule:      entryRule,
								ExitRule:       exitRule,
								}

								strategy.ShouldEnter(0, record) // returns
							}
