package main

import (
	"github.com/go-numb/go-ftx/rest/private/orders"
)

type Broker struct{}

type MainProcess struct {
	MainProcessInfo      *mainProcessInfo
	ActivePositionsCount int64
}

type mainProcessInfo struct{}

type Order struct {
	ClientID          string
	Type              string
	Market            string
	Side              string
	Price             float64
	Size              float64
	ReduceOnly        bool
	Ioc               bool
	PostOnly          bool
	RejectOnPriceBand bool
}

func (o *Order) Parse() (newOrder *orders.RequestForPlaceOrder) {
	newOrder = &orders.RequestForPlaceOrder{
		ClientID:          o.ClientID,
		Type:              o.Type,
		Market:            o.Market,
		Side:              o.Side,
		Price:             o.Price,
		Size:              o.Size,
		ReduceOnly:        o.ReduceOnly,
		Ioc:               o.Ioc,
		PostOnly:          o.PostOnly,
		RejectOnPriceBand: o.RejectOnPriceBand,
	}

	return
}

func main() {
	//login
	LoadCreds("./")

	chanOrder := make(chan Order)
	select {
	case o := <-chanOrder:

	}

}
