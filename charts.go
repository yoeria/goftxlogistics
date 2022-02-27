package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-numb/go-ftx/rest/public/markets"
)

func serveChart() {
	http.HandleFunc("/", KlineChart)
	http.ListenAndServe(":8081", nil)
}

var request, err = RestClient.Candles(&markets.RequestForCandles{ProductCode: "BTC-PERP", Resolution: 900})

type chartData struct {
	date string
	data [4]float64
}

func ParseCandles(data []markets.Candle) *[]opts.KlineData {
	items := make([]opts.KlineData, 0)
	for i := range data {
		// Shorter declaration
		kd := data[i]
		/*
			Fill data fields for iteration
			 'open', 'close', 'high', 'low'
			Fill date (time) in opts.KlineData{Name} field
		*/
		items = append(items, opts.KlineData{
			Name:  kd.StartTime.Format("2006-01-02 15:04:05"),
			Value: [4]float64{kd.Open, kd.Close, kd.High, kd.Low}})
	}

	return &items
}

func MoveTimestamp(data []opts.KlineData) (x []string) {
	for i := range data {
		x = append(x, data[i].Name)
	}
	// naked return
	return
}

func KlineChart(w http.ResponseWriter, _ *http.Request) {
	kline := charts.NewKLine()
	dataRequest := *ParseCandles(*request)

	// X axis requires a slice of strings (in this case containg the date)
	x := MoveTimestamp(dataRequest)
	// []opts.KlineData has a structure of OCLH
	y := dataRequest

	chartTitle := "Candlestick (Kline) Chart | OHLC"
	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: chartTitle,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "inside",
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
		}),
	)

	kline.SetXAxis(x).AddSeries(chartTitle, y)
	kline.Render(w)
}
