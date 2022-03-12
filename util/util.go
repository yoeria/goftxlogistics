package util

import (
	"fmt"

	"github.com/go-numb/go-ftx/rest/public/markets"
	goftxlogistics "github.com/yoeria/goftxlogistics"
)

var (
	RC = goftxlogistics.RestClient
)

// Gets the amount the account has made a trade (open-close)
func GetTradeAmount() int {
	trades, err := RC.Trades(&markets.RequestForTrades{})
	if err != nil {
		fmt.Print(err)
	}
	tradeRange := *trades
	length := len(tradeRange)
	return length
}
