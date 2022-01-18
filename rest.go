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

func LoginRest(wantMainAccount bool) *rest.Client {
	LoadCreds()
	// Only main account
	client := rest.New(auth.New(ReadonlyKey, ReadonlySecret))
	// Client with subaccounts
	clientWithSubAccounts := rest.New(
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
	// switch subaccount

	if wantMainAccount {
		return client
	} else if !wantMainAccount {
		return clientWithSubAccounts
	} else {
		return client
	}

}

func RestActions() {
	c := LoginRest(true)
	info, err := c.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account information:\n %v\n", info)

	market, err := c.Markets(&markets.RequestForMarkets{
		ProductCode: "ETH/USD",
	})

	if err != nil {
		log.Fatal(err)
	}

	// products List
	fmt.Printf("%+v\n", strings.Join(market.List(), "\n"))
	// product ranking by USD
	fmt.Printf("%+v\n", strings.Join(market.Ranking(markets.ALL), "\n"))

	// FundingRate
	rates, err := c.Rates(&futures.RequestForRates{})
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

func PlaceLimitOrder(market string, price, size float64) {
	c := LoginRest(true)

	order, err := c.PlaceOrder(&orders.RequestForPlaceOrder{
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
	c := LoginRest(true)
	// lev, err := c.Leverage(5)
	lev, err := c.Leverage(&account.RequestForLeverage{
		Leverage: 3,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", lev)

}

func GetFundingCosts() {
	c := LoginRest(true)
	funding, err := c.Funding(&funding.Request{})
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR message: %w", err))
	}
	fmt.Printf("funding: %v\n", funding)
}
