// Contains helper functions
package main

import (
	"fmt"
	"reflect"
)

// Retrieve values from chosen hash table and put those into the provided struct
func FillStrategiesStruct(s *strategies, hashTable string) error {
	var fields []string
	e := reflect.ValueOf(s).Elem()
	for i := 0; i < e.NumField(); i++ {
		fields = append(fields, e.Type().Name())
	}

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

// Retrieve values from chosen hash table and put those into the provided struct
func FillPreferencesStruct(s *Preferences, hashTable string) error {
	var fields []string
	e := reflect.ValueOf(s).Elem()
	for i := 0; i < e.NumField(); i++ {
		fields = append(fields, e.Type().Name())
	}

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
