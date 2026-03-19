package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// Client creates and verifies a Redis connection.
func Client(redisHost, redisPort string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	err := rdb.Set(context.Background(), "train:health", "ok", 0).Err()
	if err != nil {
		log.Print("redis connection failed (train-service)")
	}

	_, err = rdb.Get(context.Background(), "train:health").Result()
	if err != nil {
		log.Print("redis get check failed (train-service): " + err.Error())
	} else {
		log.Print("redis connection succeeded (train-service)")
	}

	return rdb
}
