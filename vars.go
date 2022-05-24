package main

import (
	"os"
	"sync"

	"github.com/go-numb/go-ftx/rest"
)

// All static pre-defined variables to use in the project

var (
	rc *rest.Client = RestClient
	wg sync.WaitGroup

	ReadOnlyKey    string = os.Getenv("FTX_KEY")
	ReadOnlySecret string = os.Getenv("FTX_SECRET")
	REDIS_PASSWORD string = os.Getenv("REDIS_PASSWORD")
)

const (
	STRATEGY = "strategy"
)
