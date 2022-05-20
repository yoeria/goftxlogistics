package main

import (
	"sync"

	"github.com/go-numb/go-ftx/rest"
)

// All static pre-defined variables to use in the project

var (
	rc *rest.Client = RestClient
	wg sync.WaitGroup
)

const  (
STRATEGY = "strategy"
)
