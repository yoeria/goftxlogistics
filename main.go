package main

import (
	"fmt"
)

func main() {
	//login
	LoadCreds("./")

	// Make channel that receives all Order requests
	chanOrder := make(chan Order)
	select {
	case o := <-chanOrder:
		fmt.Println(o.Side)
		break
	// Default case is check to sync config with server
	default:
		UpdateConfiguration()
		break
	}

}

// Set default values in init func
func init() {
	// Fill in rdb suppliable struct vars
	UpdateConfiguration()
}

// This function triggers the update/initialization of the configuration vars
func UpdateConfiguration() {
	activePreferences.Update()
	activeStrategies.Update()
}
