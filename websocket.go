package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-numb/go-ftx/realtime"

	"github.com/yoeria/goftxlogistics"
)

func WebsocketActions() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan realtime.Response)
	go realtime.Connect(ctx, ch, []string{"ticker"}, []string{"BTC-PERP", "ETH-PERP"}, nil)
	go LoadCreds()
	go realtime.ConnectForPrivate(ctx, ch, os.Getenv("FTX_KEY"), os.Getenv("FTX_SECRET"), []string{"orders", "fills"}, nil)

	for v := range ch {
		switch ctx.Value(v) {
		case realtime.TICKER:
			fmt.Printf("%s	%+v\n", v.Symbol, v.Ticker)

		case realtime.TRADES:
			fmt.Printf("%s	%+v\n", v.Symbol, v.Trades)
			for i := range v.Trades {
				if v.Trades[i].Liquidation {
					fmt.Printf("-----------------------------%+v\n", v.Trades[i])
				}
			}

		case realtime.ORDERBOOK:
			fmt.Printf("%s	%+v\n", v.Symbol, v.Orderbook)

		case realtime.ORDERS:
			fmt.Printf("%d	%+v\n", v.Type, v.Orders)

		case realtime.FILLS:
			fmt.Printf("%d	%+v\n", v.Type, v.Fills)

		case realtime.UNDEFINED:
			fmt.Printf("UNDEFINED %s	%s\n", v.Symbol, v.Results.Error())
		default:
			fmt.Printf("default: Errors %s	%s\n", v.Symbol, v.Results.Error())

		}

	}
}
