package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/gookit/color"
	"github.com/labstack/gommon/color"
)

// generate random data for line chart
func generateKlineItems() []opts.KlineData {
	items := make([]opts.KlineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.KlineData{Value: rand.Intn(300)})
	}
	return items
}

func httpServer(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func serveChart() {
	http.HandleFunc("/", httpServer)
	http.ListenAndServe(":8081", nil)
}

var request, err = RestClient.Candles(&markets.RequestForCandles{ProductCode: "BTC-PERP", Resolution: 900})
func parseCandles(data markets.ResponseForCandles) ([]chartData){
/* 	candle := &markets.Candle{StartTime: *data[0]}
 */
	parsedContent := make([]chartData,0)
 	for iteration := range data {
		parsedContent = append(parsedContent, data[iteration])

	fmt.Println(data[iteration])
	fmt.Println("Iteration number", color.Magenta.Render((iteration))  )
	}
	return parsedContent
}

type chartData struct {
	date string `time.Time`
	data [4]float64
}

func TestReq(t *testing.T) {
	fmt.Printf("t: %v\n", t)
	fmt.Println(request)
	fmt.Println()
}

/* var kd = make([...]opts.KlineData{
	Name: "BTC-PERP",Value:
},0)
func klineChart(w http.ResponseWriter, _ *http.Request) *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, opts.KlineData{Value: kd[i].data})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "DataZoom(inside&slider)",
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

	kline.SetXAxis(x).AddSeries("kline", y)
	return kline
}
*/
