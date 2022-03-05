package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-numb/go-ftx/rest/public/futures"
)

var (
	rc = RestClient
)

func currentPotentFutures() (list []futures.FutureForList) {
	getFutures, err := rc.Futures(&futures.RequestForFutures{})
	if err != nil {
		return
	}
	for range *getFutures {
		list = append(list, *getFutures...)
	}
	spew.Dump(getFutures)
	log.Println(len(list))
	return list
}
