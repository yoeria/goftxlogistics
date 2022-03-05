package main

import (
	"fmt"
	//"log"

	"sort"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
	"github.com/go-numb/go-ftx/rest/private/funding"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/rest/public/futures"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/go-numb/go-ftx/types"
	"github.com/labstack/gommon/log"
)

var (
	RestClient = rest.New(auth.New(ReadonlyKey, ReadonlySecret))
	ClientWithSubAccounts = rest.New(
	auth.New(
		ReadonlyKey,
		ReadonlySecret,
		auth.SubAccount{
			UUID:     1,
			Nickname: "MM",
		},
		auth.SubAccount{
			UUID:     2,
			Nickname: "Test",
		},
	))
)

func RestActions() {
	info := getAccountInfo()

	fmt.Printf("Account information:\n %v\n", info)

	market := getMarkets()

	// products List
	fmt.Printf("%+v\n", strings.Join(market.List(), "\n"))
	// product ranking by USD
	fmt.Printf("%+v\n", strings.Join(market.Ranking(markets.ALL), "\n"))

	// FundingRate
	rates, err := RestClient.Rates(&futures.RequestForRates{})
	if err != nil {
		log.Fatal(err)
	}
	// Sort by FundingRate & Print
	// Custom sort
	sort.Sort(sort.Reverse(rates))
	for _, v := range *rates {
		fmt.Printf("%s			%s		%s\n", humanize.Commaf(v.Rate), v.Future, v.Time.String())
	}
}

func getMarkets() *markets.ResponseForMarkets {
	market, err := RestClient.Markets(&markets.RequestForMarkets{
	})

	if err != nil {
		log.Fatal(err)
	}
	return market
}

func getAccountInfo() *account.ResponseForInformation {
	info, err := RestClient.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("Account information:\n %v\n", info)
	return info
}

func PlaceLimitOrder(market string, price, size float64) {

	order, err := RestClient.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.LIMIT,
		Market: market,
		Side:   types.BUY,
		Price:  price,
		Size:   size,
		// Optionals
		ClientID:   "use_original_client_id",
		Ioc:        false,
		ReduceOnly: false,
		PostOnly:   false,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", order)
}

func SetAccountLeverage() {

	// lev, err := RestClient.Leverage(5)
	lev, err := RestClient.Leverage(&account.RequestForLeverage{
		Leverage: 3,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", lev)

}

func GetFundingCosts() {
	funding, err := RestClient.Funding(&funding.Request{})
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR message: %w", err))
	}
	fmt.Printf("funding: %v\n", funding)
}
