package main

import (
	"fmt"
	"log"

	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
	"github.com/go-numb/go-ftx/rest/public/markets"
)

func getAccountInformation(rc *rest.Client) *account.ResponseForInformation {
	info, err := rc.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}
	return info
}

// Gets the amount the account has made a trade (open-close)
func GetTradeAmount(rc *rest.Client) int {
	trades, err := rc.Trades(&markets.RequestForTrades{})
	if err != nil {
		fmt.Print(err)
	}
	tradeRange := *trades
	length := len(tradeRange)
	return length
}

// Gets the count of the account's current positions '
func GetCurrentPositionCount(rc *rest.Client) int {
	count := len(getAccountInformation(rc).Positions)
	return count
}
