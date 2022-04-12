package account

import (
	"os"
	"sort"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-numb/go-ftx/rest/private/fills"
	"github.com/yoeria/goftxlogistics/rest"
)

func TestGetActiveTradeFundingCosts(t *testing.T) {
	rest.LoadCreds("../rest/creds.env")
	rc := rest.ClientWithSubAccounts
	// Use MM Subaccount
	rc.Auth.UseSubAccountID(1)
	rc.Auth.Key = os.Getenv("FTX_KEY")
	rc.Auth.Secret = os.Getenv("FTX_SECRET")
	reqFills, err := rc.Fills(&fills.Request{})
	if err != nil {
		return
	}
	getFills := *reqFills
	var fillSlice fills.Fill
	var fillTimeSlices []time.Time
	for range getFills {
		fillTimeSlices = append(fillTimeSlices, fillSlice.Time)
	}
	sort.Slice(fillTimeSlices, func(i, j int) bool { return fillTimeSlices[i].Unix() < fillTimeSlices[j].Unix() })
	spew.Dump(fillTimeSlices)
}
