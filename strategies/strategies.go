package strategies

import (
	"github.com/sdcoffey/techan"
)

// Indicator average ratio of at least 0.8 (80%) should start buying cycle
// If calculateSuccessRatio >= 0.8 -> trigger an order
func calculateSuccessRatio(_ chan int) {

	for {
		select {}
	}

}

// Calculate FAST StochasticRSI values
func calcFastStochRSI(series *techan.TimeSeries, timeframe int) {
	techan.NewFastStochasticIndicator(series, timeframe)
}
