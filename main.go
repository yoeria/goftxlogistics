package main

import (
	"github.com/go-numb/go-ftx/rest/private/orders"
)

type Broker struct{}

type MainProcess struct {
	MainProcessInfo      *mainProcessInfo
	ActivePositionsCount int64
	Action               chan Buy
}

type mainProcessInfo struct{}

// Buy instruction details
type Buy chan struct {
	Currency string
}

// Sell instruction details
type Sell chan struct {
	Currency string
}
type Order *orders.RequestForPlaceOrder
func main() {
	//login
	LoadCreds("./")
	Order
	chanBuy := make(chan Buy)
	chanSell := make(chan Sell)
	select {
	case b := <-chanBuy:
		rc.PlaceOrder(&orders.RequestForPlaceOrder{
			ClientID:          "use_original_client_id",
			Type:              "buy",
			Market:            ,
			Side:              "",
			Price:             0,
			Size:              0,
			ReduceOnly:        false,
			Ioc:               false,
			PostOnly:          false,
			RejectOnPriceBand: false,
		})
	case s <- chanSell:
		rc.PlaceOrder(Order)
	}

}
