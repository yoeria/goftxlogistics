package strategies

import (
	talib "github.com/markcheno/go-talib"
	"github.com/sdcoffey/techan"
)

// Indicator average ratio of at least 0.8 (80%) should start buying cycle
// If calculateSuccessRatio >= 0.8 -> trigger an order
func calculateSuccessRatio(_ chan int) {

	for {
		select {}
	}

}

// Calculate FAST Stochastic values
func calcFastStoch(series *techan.TimeSeries, timeframe int) {
	techan.NewFastStochasticIndicator(series, timeframe)

}

func calcStochRSI(inReal []float64, timeframe, kPeriod, dPeriod int, maType talib.MaType) ([]float64, []float64) {
	return talib.StochRsi(inReal, timeframe, kPeriod, dPeriod, maType)
}

func calcEMA() {
	//techan.NewEMAIndicator()
}
