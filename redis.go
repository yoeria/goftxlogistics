package main

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     "redis.dssq.xyz",
	Password: "coliseum-dealer-entree",
	DB:       0, // use default DB
})

// Test validity of redis connection.
func Test(t *testing.T) {
	testCommand := RDB.Info(RDB.Context())

	fmt.Printf("t: %v\n", t)
	fmt.Println(testCommand)
}
