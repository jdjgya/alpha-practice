package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "0.0.0.0:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)

	err := rdb.Set("key", "test_value", time.Second*60).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
