package main

import (
	"fmt"

	cmd "github.com/yoeria/goftxlogistics/cmd"
	"github.com/yoeria/goftxlogistics/pocketbase"
)

// Set default values in init func
func init() {
	// Fill in rdb suppliable struct vars
	UpdateConfiguration()
	go cmd.Exec()
	wg.Add(1)
	go pocketbase.Exec()
}

// Launch instances by creating a variable with the value of struct ProcessInformation
func main() {

	// Make instance of ProcessInformation
	p0 := &ProcessInformation{
		// Each process is keeping their own statistics
		Statistics: &Statistics{
			ActivePositionsCount: 0,
		},
		Preferences: aPreferences,

		SyncInfo: aSyncInfo,
	}

	// Make channel that receives all Order requests
	ChanOrder := make(chan Order)
	select {
	case order := <-ChanOrder:
		if aPreferences.MaxConcurrentPositions > p0.Statistics.ActivePositionsCount {
			order.Exec()
			// TODO: Place check routine if order is filled or not
		}
		fmt.Println(order.Side)
		break
	// Default case is check to sync config with server
	default:
		UpdateConfiguration()
	}

	// Block ending of main process until waitgroup is empty
	wg.Wait()
}

// This function triggers the update/initialization of the configuration vars
func UpdateConfiguration() {
	go aPreferences.Update()
	go aStrategies.Update()
}
