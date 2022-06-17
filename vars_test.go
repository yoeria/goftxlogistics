package main

import (
	"testing"
)

// Test validity of redis connection.
func Test(t *testing.T) {
	ctx := rdb.Context()
	_, err := rdb.Info(ctx).Bool()
	testName := "Redis Connection"
	if err != nil {
		t.Errorf("%vREASON: %v", getTestConclusion(false, testName), err)
	}
	getTestConclusion(true, testName)

}
