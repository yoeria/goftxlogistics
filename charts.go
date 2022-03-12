package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/yoeria/goftxlogistics/rest"
)

func serveChart() {
	http.HandleFunc("/", KlineChart)
	http.ListenAndServe(":8081", nil)
}

var request, err = rest.RestClient.Candles(&markets.RequestForCandles{ProductCode: "BTC-PERP", Resolution: 900})

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
/* Kline example:
https://github.com/go-echarts/examples/blob/master/examples/kline.go */
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
	)

	kline.SetXAxis(x).AddSeries(chartTitle, y).SetSeriesOptions(
			charts.WithMarkPointNameTypeItemOpts(opts.MarkPointNameTypeItem{
				Name:     "highest value",
				Type:     "max",
				ValueDim: "highest",
			}),
			charts.WithMarkPointNameTypeItemOpts(opts.MarkPointNameTypeItem{
				Name:     "lowest value",
				Type:     "min",
				ValueDim: "lowest",
			}),
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show: true,
				},
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color:        "#ec0000",
				Color0:       "#00da3c",
				BorderColor:  "#8A0000",
				BorderColor0: "#008F28",
			}),
		)
	kline.Render(w)
}
