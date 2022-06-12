package main

import (
	"fmt"
	"testing"
)

// Test validity of redis connection.
func Test(t *testing.T) {
	LoadCreds("./")
	ctx := rdb.Context()
	testCommand, err := rdb.Info(ctx).Result()
	if err != nil {
		t.Errorf("❌ Redis Connection\n REASON: %v", err)
	}
	fmt.Println(testCommand)

}
