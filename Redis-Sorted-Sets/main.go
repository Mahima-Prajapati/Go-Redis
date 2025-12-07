package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go + Redis + SortedSets")

	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		DB: 3,
	})

	// Flush FIRST so DB starts empty every run
	rdb.FlushDB(ctx)

	// { ZAdd } method example
	resInt, _ := rdb.ZAdd(ctx, "key",
		redis.Z{4, "filed4"},
		redis.Z{2, "field2"},
		redis.Z{1, "field1"},
		redis.Z{3, "field3"},
		redis.Z{4, "filed4"}).Result()
	fmt.Println(resInt) // ----> 4

	// { ZRank } method example
	resInt, _ = rdb.ZRank(ctx, "key", "field1").Result()
	fmt.Println(resInt) // ----> 0

	resInt, _ = rdb.ZRank(ctx, "key", "field3").Result()
	fmt.Println(resInt) // ----> 2

	// { ZRevRank } method example
	resInt, _ = rdb.ZRevRank(ctx, "key", "field1").Result()
	fmt.Println(resInt) // ----> 3

	resInt, _ = rdb.ZRevRank(ctx, "key", "field3").Result()
	fmt.Println(resInt) // -----> 1

	// { ZRange } method example
	resStrList, _ := rdb.ZRange(ctx, "key", 1, 3).Result()
	fmt.Println(resStrList) // [field2 field3 filed4]

	// { ZRevRange } method example
	resStrList, _ = rdb.ZRevRange(ctx, "key", 2, 3).Result()
	fmt.Println(resStrList) // [field2 field1]
}
