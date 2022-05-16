package main

import (
	"github.com/go-redis/redis/v8"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     "dssq.xyz:6380",
	Password: "coliseum-dealer-entree",
	DB:       0, // use default DB
})
