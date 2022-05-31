package main

import (
	"net/http"
	"testing"
)

/* func TestReq(t *testing.T) {
	fmt.Printf("t: %v\n", t)
	color.Set(color.HiGreen, color.BgBlack)
	spew.Dump(*ParseCandles(*request))
	color.Reset()
}

func TestChart(t *testing.T) {
	fmt.Print("TestChart func:\n\n")
	color.Set(color.Magenta)
	color.Reset()
	serveChart()
} */

func Test_KlineChart(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			KlineChart(tt.args.w, tt.args.in1)
		})
	}
}
