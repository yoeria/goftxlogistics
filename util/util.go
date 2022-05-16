package util

import (
	"fmt"

	"github.com/go-numb/go-ftx/rest/public/markets"
)

var (
	rc = RestClient
)

// Gets the amount the account has made a trade (open-close)
func GetTradeAmount() int {
	trades, err := rc.Trades(&markets.RequestForTrades{})
	if err != nil {
		fmt.Print(err)
	}
	tradeRange := *trades
	length := len(tradeRange)
	return length
}
