package main

import (
	"fmt"
	//"log"
	"os"
	"sort"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/go-numb/go-ftx/auth"
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
	"github.com/go-numb/go-ftx/rest/private/orders"
	"github.com/go-numb/go-ftx/rest/public/futures"
	"github.com/go-numb/go-ftx/rest/public/markets"
	"github.com/go-numb/go-ftx/types"
	"github.com/labstack/gommon/log"
)

func LoginRest(wantMainAccount bool) *rest.Client {
	LoadCreds()
	readonlyKey := os.Getenv("FTX_KEY")
	readonlySecret := os.Getenv("FTX_SECRET")
	// Only main account
	client := rest.New(auth.New(readonlyKey, readonlySecret))
	// Client with subaccounts
	clientWithSubAccounts := rest.New(
		auth.New(
			readonlyKey,
			readonlySecret,
			auth.SubAccount{
				UUID:     1,
				Nickname: "MM",
			},
			auth.SubAccount{
				UUID:     2,
				Nickname: "Test",
			},
			// many....
		))
	// switch subaccount
	clientWithSubAccounts.Auth.UseSubAccountID(1)

	return client

}

func RestActions() {
	// Creds

	// or
	// UseSubAccounts
	// or 2... this number is key in map[int]SubAccount

	// account informations
	// client or clientWithSubAccounts in this time.
	c := LoginRest(true) // or clientWithSubAccounts
	info, err := c.Information(&account.RequestForInformation{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account information:\n %v\n", info)

	market, err := c.Markets(&markets.RequestForMarkets{
		ProductCode: "XRPBULL/USDT",
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

	order, err := c.PlaceOrder(&orders.RequestForPlaceOrder{
		Type:   types.LIMIT,
		Market: "BTC-PERP",
		Side:   types.BUY,
		Price:  6200,
		Size:   0.01,
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

	ok, err := c.CancelByID(&orders.RequestForCancelByID{
		OrderID: 1,
		// either... , prioritize clientID
		ClientID:       "",
		TriggerOrderID: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok)
	// ok is status comment

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
