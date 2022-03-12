package util

import (
	"fmt"

	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/yoeria/goftxlogistics/rest"
)

var (
	RC = rest.RestClient
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
