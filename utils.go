// Contains helper functions
package main

import (
	"fmt"
	"reflect"
)

// Retrieve values from chosen hash table and put those into the provided struct
func UpdateStrategiesStruct(s *strategies, hashTable string) error {
	fields := getStrategyFields(s)
	getValues := rdb.HMGet(ctx, hashTable, fields...)
	values, err := getValues.Result()
	if err != nil {
		return err
	}

	// TODO
	// Put received values into the provided struct
	for k, v := range values {
		fmt.Printf("Key:\t%v\nValue:\t%v", k, v)

	}

	// In case of no error
	return nil
}

// Since the fields could be updated in code the fields are dynamic. Reflect get the structs' fieldnames
func getStrategyFields(s *strategies) (fields []string) {
	val := reflect.ValueOf(s).Elem()

	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//tag := typeField.Tag

		fields = append(fields, typeField.Name)
	}
	return
}

// Retrieve values from chosen hash table and put those into the provided struct
func UpdatePreferencesStruct(p *Preferences, hashTable string) error {

	fields := getPreferencesFields(p)

	getValues := rdb.Pipeline().HMGet(ctx, hashTable, fields...)
	values, err := getValues.Result()
	if err != nil {
		return err
	}

	// TODO
	// Update the struct with the received values from db
	for i := range values {
		intValues, ok := (values)[i].(int64)
		if !ok {
			fmt.Println("Converting from []interface{} to int64 failed")
			panic(1)
		}

		(*p).MaxConcurrentPositions = intValues

	}

	// In case of no error
	return nil
}

// Since the fields could be updated in code the fields are dynamic. Reflect get the structs' fieldnames
func getPreferencesFields(p *Preferences) (fields []string) {
	val := reflect.ValueOf(p).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		// Checks if value of field is a pointer; If so, continue
		if valueField.Type().Kind() == reflect.Ptr {
			continue
		}

		fields = append(fields, typeField.Name)
	}
	return
}

/* // TODO combine a rest request of historic candle data with a websocket request for the same asset/crypto
func CandleStream(restReq *markets.RequestForCandles)(chTimeSeries chan *techan.TimeSeries){
	// Create variables for receiving the request data
	restData,err:= rc.Candles(restReq)
	wsData:= make(chan realtime.Response)
    // Connect to websocket stream (ticker)
	go realtime.Connect(ctx, wsData, []string{"ticker"}, []string{restReq.ProductCode}, nil)
	select {
	case ch := <-wsData:

	}


}
*/

// TODO func for quick indicator followup (second timeframes), make order squeezer (between ask/bid)
