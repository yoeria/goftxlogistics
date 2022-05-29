package main

import (
	"encoding/json"
	"fmt"
)

// Configuration for XYZ
type Preferences struct {
	MaxConcurrentPositions int64
	EnabledStrategies      *strategies
}

type strategies struct {
	EMA      bool
	SMA      bool
	STOCHRSI bool
}

func (s *strategies) Update() {
	hashkey := "strategiesConfig" + ":" + string(HASHES["strategies"])
	values := rdb.HMGet(ctx,hashkey,"*")
	for k, v := range values.Val() {
		fmt.Printf("Key:\t%v\nValue:\t%v",k,v)

	}
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
