package main

import (
	"context"
	"os"
	"sync"
	"time"

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

	// Time config was last synced on local machine
	timeOfLastSync time.Time
	// Bool is true if changes are made locally and they need to be synced with redis server
	localChangesDone bool
	// Time config was last changed locally
	timeOfLocalChanges time.Time
	// Time config was last changed on the redis server
	timeOfOnlineChanges time.Time

	HASHES map[string]int64 = map[string]int64{
		"preferences":0,
		"strategies":1,
	}

)

const (
	STRATEGY = "strategy"
)
