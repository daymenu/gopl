package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	ExampleNewClient()
}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("name", "hanjian", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)
	// Output: PONG <nil>
}
