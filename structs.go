package main

import (
	json "github.com/json-iterator/go"
	"fmt"
	"log"
	"time"

	"github.com/go-numb/go-ftx/rest/private/orders"
)

// TODO: Add Log field containing a logger from "log" package
type ProcessInformation struct {
	Preferences *Preferences `json:"preferences"`
	Statistics  *Statistics  `json:"statistics"`
	SyncInfo    *SyncInfo    `json:"syncinfo"`
}

// Configuration for a process
type Preferences struct {
	MaxConcurrentPositions int64
	EnabledStrategies      *strategies
}

// Update/init the struct
func (p *Preferences) Update() error {
	return FillPreferencesStruct(p, "preferences")
}

type strategies struct {
	EMA      bool
	SMA      bool
	STOCHRSI bool
}

// Update/init the struct
func (s *strategies) Update() error {
	return FillStrategiesStruct(s, "strategies")
}

type Statistics struct {
	ActivePositionsCount int64
}

type SyncInfo struct {
	// Time config was last synced on local machine
	timeOfLastSync time.Time
	// Bool is true if changes are made locally and they need to be synced with redis server
	localChangesDone bool
	// Time config was last changed locally
	timeOfLocalChanges time.Time
	// Time config was last changed on the redis server
	timeOfOnlineChanges time.Time
}

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
