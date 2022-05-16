package main

import (
	"context"
	"fmt"

	"github.com/go-numb/go-ftx/realtime"
)

func WebsocketActions() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan realtime.Response)
	go realtime.Connect(ctx, ch, []string{"ticker"}, []string{"BTC-PERP"}, nil)
	go LoadCreds("./rest/creds.env")
	go realtime.ConnectForPrivate(ctx, ch, ReadOnlyKey, ReadOnlySecret, []string{"orders", "fills"}, nil)

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
