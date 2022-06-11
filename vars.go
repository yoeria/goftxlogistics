package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/go-numb/go-ftx/rest"
	"github.com/go-redis/redis/v8"
)

// All STATIC pre-defined variables to use in the project

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

	HASHES map[string]int64 = map[string]int64{
		"preferences": 0,
		"strategies":  1,
	}
)

// All DYNAMIC pre-defined variables to use in the project
var (
	// Use the aXXXXXXX variables for every instance since these values have global properties and are not instance specific

	// Reading values from redis
	aPreferences *Preferences = &Preferences{
		MaxConcurrentPositions: 0,
		EnabledStrategies:      aStrategies,
	}
	aStrategies *strategies = &strategies{
		EMA:      false,
		SMA:      false,
		STOCHRSI: false,
	}
	// Reading and setting values locally
	aSyncInfo *SyncInfo = &SyncInfo{
		timeOfLastSync:      time.Time{},
		localChangesDone:    false,
		timeOfLocalChanges:  time.Time{},
		timeOfOnlineChanges: time.Time{},
	}
)

const (
	STRATEGY = "strategy"
)
