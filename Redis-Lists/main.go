package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go + Redis + Lists")
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	// Flush FIRST so DB starts empty every run
	rdb.FlushDB(ctx)

	// { LPush } method example
	resInt, _ := rdb.LPush(ctx, "key1", 1, 2).Result()
	fmt.Println(resInt) // ----> 2

	resInt, _ = rdb.LPush(ctx, "key1", 3, 4).Result()
	fmt.Println(resInt) // -----> 4

	// { RPush } method example
	resInt, _ = rdb.RPush(ctx, "key1", 5, 6).Result()
	fmt.Println(resInt) // ----> 6

	resInt, _ = rdb.RPush(ctx, "key1", 7, 8).Result()
	fmt.Println(resInt) // -----> 8

	// { LPop } method example
	resStr, _ := rdb.LPop(ctx, "key1").Result()
	fmt.Println(resStr) // -----> 4

	// { RPop } method example
	resStr, _ = rdb.RPop(ctx, "key1").Result()
	fmt.Println(resStr) // -----> 8

	// { LLen } method example
	resInt, _ = rdb.LLen(ctx, "key1").Result()
	fmt.Println(resInt) // -----> 6

	// { LRange } method example
	res, _ := rdb.LRange(ctx, "key1", 1, 4).Result()
	fmt.Println(res) // -----> [2, 1, 5, 6] inclusive

	// { LTrim } method example
	resStr, _ = rdb.LTrim(ctx, "key1", 1, 3).Result()
	fmt.Println(resStr) // ----> OK

	// { LMove } method example
	resStr, _ = rdb.LMove(ctx, "key1", "key2", "LEFT", "LEFT").Result()
	fmt.Println(resStr) // -----> 2

	resStr, _ = rdb.LMove(ctx, "key1", "key2", "RIGHT", "LEFT").Result()
	fmt.Println(resStr) // -----> 5
}
