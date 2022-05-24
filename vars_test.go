package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/gookit/color"
)

// Test validity of redis connection.
func Test(t *testing.T) {
	ctx := rdb.Context()
	testCommand := rdb.Info(ctx)

	fmt.Printf("t: %v\n", t)
	fmt.Println(testCommand)
}

func TestAppendTradeToTrades(t *testing.T) {
	ctx := rdb.Context()
	count := 100
	for i := 0; i < count; i++ {
		byteSlice, err := Trade{
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
		add := rdb.ZAdd(ctx, "TradesTesting", &redis.Z{
			Score:  rand.Float64(),
			Member: byteSlice,
		})

		result, err := add.Result()
		fmt.Println(string(rune(result)))
		if result != 0 {
			color.Errorln(err)
			t.Error()
		}

		if add.Err() != nil {
			color.Redln("error triggered:")
			fmt.Println(add.Err().Error())
			t.Error()
		}

	}

}
