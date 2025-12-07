package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default port
		Password: "",               // No password set
	})

	
}
