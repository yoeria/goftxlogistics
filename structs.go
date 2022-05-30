package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-numb/go-ftx/rest/private/orders"
)

// Configuration for XYZ
type Preferences struct {
	MaxConcurrentPositions int64
	EnabledStrategies      *strategies
}

func (p *Preferences) Update() error {
	return FillStruct(p,"preferences")
}

type strategies struct {
	EMA      bool
	SMA      bool
	STOCHRSI bool
}

// Update/init the struct
func (s *strategies) Update() error {
	return FillStruct(s, "strategies")
}

type Statistics struct {
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

// Executes sending (trying to place) order to server
func (o *Order) Exec() (order *orders.ResponseForPlaceOrder, err error) {
	order, err = rc.PlaceOrder(o.Parse())
	if err != nil {
		log.Println(err)
	}
	// logging orderstatus in stdout
	fmt.Println(order.Status)
	return
}

//Used to keep track of trades made by the system
type Trade struct {
	// Date and time on entering trade
	DatetimeEnter string `json:"datetime_enter"`
	// Date and time on exiting trade
	DatetimeExit string `json:"datetime_exit"`
	// Price of asset on entering trade
	Enter float64 `json:"price_enter"`
	// Price of asset when exiting trade
	Exit float64 `json:"price_exit"`
	// The net result in the unit that was traded
	Net float64 `json:"roe_unit"`
	// The net result of the trade in percent
	NetPercent float64 `json:"roe_percent"`
	// Value (USD) of overall wallet after closing of this trade
	Wallet float64 `json:"wallet_exit"`
	// Net impact on wallet value in percentage
	WalletNet float64 `json:"wallet_net_percentage_exit"`
	// Total costs to platform for trade (funding included if applicable)
	Fees float64 `json:"fees"`
}

func (t Trade) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t Trade) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &t)
}
