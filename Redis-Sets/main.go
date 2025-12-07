package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go + Redis + Sets")

	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		DB: 2,
	})

	// Flush FIRST so DB starts empty every run
	rdb.FlushDB(ctx)

	// { SAdd } method example
	resInt, _ := rdb.SAdd(ctx, "key1", 1, 2, 3, 4, 1).Result()
	fmt.Println(resInt) // ----> 4

	resInt, _ = rdb.SAdd(ctx, "key1", 2).Result()
	fmt.Println(resInt) // ----> 0

	// { SPop } method example
	resStr, _ := rdb.SPop(ctx, "key1").Result()
	fmt.Println(resStr) // -----> 2 (random-number from set)

	// { SIsMember } method example
	resBool, _ := rdb.SIsMember(ctx, "key1", 3).Result()
	fmt.Println(resBool) // ----> true

	resBool, _ = rdb.SIsMember(ctx, "key1", 5).Result()
	fmt.Println(resBool) // ----> false

	// { SRem } method example
	resInt, _ = rdb.SRem(ctx, "key1", 3, 4).Result()
	fmt.Println(resInt) // ----> 2

	resInt, _ = rdb.SCard(ctx, "key1").Result()
	fmt.Println(resInt) // ----> 1
}
