package db

import (
	"fmt"
	"testing"
)

// Test validity of redis connection.
func Test(t *testing.T) {
	ctx := RDB.Context()
	testCommand := RDB.Info(ctx)

	fmt.Printf("t: %v\n", t)
	fmt.Println(testCommand)
}
