package strategies

import "github.com/sdcoffey/techan"

// Format for delivering the required data for calculating the indicators.
type StratData struct {
	timeFrame int
	techanSeries *techan.TimeSeries
	
}
