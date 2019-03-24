package model

import (
	"github.com/go-redis/redis"
)

var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
