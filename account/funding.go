package account

import (
	"github.com/go-numb/go-ftx/rest"
	"github.com/go-numb/go-ftx/rest/private/account"
)

func GetActiveTradeFundingCosts(rc *rest.Client) (fundingCosts float64) {
	rc.Positions(&account.RequestForPositions{})
	return
}
