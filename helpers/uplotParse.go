<<<<<<< HEAD
type OhlcObject struct {
	Result  []StOHLCv `json:"result"`
	Success bool      `json:"success"`
}

type StOHLCv struct {
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	StartTime string  `json:"startTime"`
	Time      float64 `json:"time"`
	Volume    float64 `json:"volume"`
}

type FundingPaymentsObject struct {
	Result  []FundingPaymentsResult `json:"result"`
	Success bool                    `json:"success"`
}

type FundingPaymentsResult struct {
	Future  string  `json:"future"`
	ID      float64 `json:"id"`
	Payment float64 `json:"payment"`
	Rate    float64 `json:"rate"`
	Time    string  `json:"time"`
}
=======
package main

func parseData([]string) {

}
>>>>>>> 24392a3 (commit for usage on laptop)
