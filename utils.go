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
	// Put received values into the provided struct
	for k, v := range values {
		fmt.Printf("Key:\t%v\nValue:\t%v", k, v)

	}

	// In case of no error
	return nil
}

func getPreferencesFields(p *Preferences) (fields []string) {
	val := reflect.ValueOf(p).Elem()

	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//tag := typeField.Tag

		fields = append(fields, typeField.Name)
	}
	return
}

// Pass if test is a successful test and the name of the test
func getTestConclusion(isSuccess bool, testName string) string {
	// if TRUE
	if isSuccess {
		return fmt.Sprintf("✅\t%v\n", testName)
	}
	// if FALSE
	return fmt.Sprintf("❌\t%v\n", testName)
}

func getReason(err error) string {
	return fmt.Sprintf("❓REASON:\n%v", err)
}
