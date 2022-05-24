package main

import (
	"context"
	"os"
	"sync"

	"github.com/go-numb/go-ftx/rest"
	"github.com/go-redis/redis/v8"
)

// All static pre-defined variables to use in the project

var (
	rc  *rest.Client  = RestClient
	rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "dssq.xyz:6380",
		Password: REDIS_PASSWORD,
		DB:       0, // use default DB
	})
	wg  sync.WaitGroup
	ctx context.Context = context.Background()

	ReadOnlyKey    string = os.Getenv("FTX_KEY")
	ReadOnlySecret string = os.Getenv("FTX_SECRET")
	REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")
)

const (
	STRATEGY = "strategy"
)
