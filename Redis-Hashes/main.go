package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		DB: 3,
	})

	// Flush FIRST so DB starts empty every run
	rdb.FlushDB(ctx)

	// { HSet } method example
	resInt, _ := rdb.HSet(ctx, "key", "field1", 1, "field3", 3).Result()
	fmt.Println(resInt) // ----> 2

	resInt, _ = rdb.HSet(ctx, "key", "field2", 2).Result()
	fmt.Println(resInt) // ----> 1

	// { HGet } method example
	resStr, _ := rdb.HGet(ctx, "key", "field1").Result()
	fmt.Println(resStr) // -----> 1

	// { HKeys } method example
	resIntList, _ := rdb.HKeys(ctx, "key").Result()
	fmt.Println(resIntList) // [field1 field3 field2]

	// { HVals } method example
	resStrList, _ := rdb.HVals(ctx, "key").Result()
	fmt.Println(resStrList) // [1 3 2]

	// { HGetAll } method example
	resList, _ := rdb.HGetAll(ctx, "key").Result()
	fmt.Println(resList) // map[field1:1 field2:2 field3:3]

	// { HIncrBy } method example
	resInt, _ = rdb.HSet(ctx, "hash", "count", 1).Result()
	fmt.Println(resInt) // -----> 1

	resInt, _ = rdb.HIncrBy(ctx, "hash", "count", 4).Result()
	fmt.Println(resInt) // -----> 5
}
