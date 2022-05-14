package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"github.com/yoeria/goftxlogistics/rest"
)

func main() {
	rest.LoadCreds("./rest/creds.env")
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

func main() {
	//login
	rest.LoadCreds("./rest/creds.env")


	

}
