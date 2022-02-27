package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
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

func ParseCandles(data markets.ResponseForCandles) (*[]chartData, *[]opts.KlineData) {
	parsedData := make([]chartData, 0)
	items := make([]opts.KlineData, 0)
	for i := 0; i < data.Len(); i++ {
		// Shorter declaration
		kd := data[i]
		items = append(items, opts.KlineData{Value: [4]float64{kd.Open, kd.Close, kd.High, kd.Low}})
		// Fill date (time) field for iteration
		parsedData[i].date = kd.StartTime.Format("2006-01-02 15:04:05")
		// Fill data fields for iteration
		// 'open', 'close', 'high', 'low'
		parsedData[i].data = [4]float64{kd.Open, kd.Close, kd.High, kd.Low}
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

	// X axis requires a slice of strings
	x := make([]string, 0)
	// []opts.KlineData has a structure of OCLH
	y := make([]opts.KlineData, 0)

	usableData := *ParseCandles(*request)
	spew.Dump(usableData)
	for i := range usableData {
		y[i].Value = usableData[i]
		x[i] = usableData[i].date
	}

	// Logging purposes
	//spew.Dump(x,y)

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
	)

	kline.SetXAxis(x).AddSeries(chartTitle, y)
	kline.Render(w)
}
