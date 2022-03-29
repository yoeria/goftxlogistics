package structs

import (
	"encoding/json"
)

/*
Used to keep track of trades made by the system

Keep track in Redis DB
*/
type Trade struct {
	// Date and time on entering trade
	DatetimeEnter string ``
	// Date and time on exiting trade
	DatetimeExit string
	// Price of asset on entering trade
	Enter float64
	// Price of asset when exiting trade
	Exit float64
	// The net result in the unit that was traded
	Net float64
	// The net result of the trade in percent
	NetPercent float64
	// Value (USD) of overall wallet after closing of this trade
	Wallet float64
	// Net impact on wallet value in percentage
	WalletNet float64
	// Total costs to platform for trade (funding included if applicable)
	Fees float64
}

func (t Trade) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t Trade) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &t)
}
