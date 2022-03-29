package db

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/color"
	"github.com/yoeria/goftxlogistics/structs"
)

// Test validity of redis connection.
func Test(t *testing.T) {
	ctx := RDB.Context()
	testCommand := RDB.Info(ctx)

	fmt.Printf("t: %v\n", t)
	fmt.Println(testCommand)
}

func TestAppendTradeToTrades(t *testing.T) {
	ctx := RDB.Context()
	count := 100
	for i := 0; i < count; i++ {
		byteSlice, err := structs.Trade{
			DatetimeEnter: "2022-01-01",
			DatetimeExit:  "2022-01-02",
			Enter:         float64(i ^ 2),
			Exit:          float64(i ^ 2),
			Net:           float64(i ^ 2),
			NetPercent:    float64(i ^ 2),
			Wallet:        float64(i ^ 2),
			WalletNet:     float64(i ^ 2),
			Fees:          float64(i ^ 2),
		}.MarshalBinary()
		if err != nil {
			fmt.Println(err.Error())
		}
		add := RDB.ZAdd(ctx, "TradesTesting", &redis.Z{
			Score:  rand.Float64(),
			Member: byteSlice,
		})
		if add.Err().Error() != "" {
			color.Redln("error triggered:")
			fmt.Println(add.Err().Error())
			t.Error()
		}

	}

}
