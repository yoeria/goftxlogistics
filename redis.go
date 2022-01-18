package main

import (
	"github.com/go-redis/redis"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:	  "redis.dssq.xyz",
	Password: "coliseum-dealer-entree",
	DB:		  0,  // use default DB
})
