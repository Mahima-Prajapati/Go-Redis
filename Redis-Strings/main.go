package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Go + Redis + Strings")

	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default port
		Password: "",               // No password set
		DB:       0,
	})

	// Flush FIRST so DB starts empty every run
	rdb.FlushDB(ctx)

	// { Set } method example
	resStr, _ := rdb.Set(ctx, "key1", "value1", 0).Result()
	fmt.Println(resStr) // ----> OK

	resStr, _ = rdb.Set(ctx, "key2", "value2", 0).Result()
	fmt.Println(resStr) // ----> OK

	resStr, _ = rdb.Set(ctx, "key1", "older-value", 0).Result()
	fmt.Println(resStr) // ----> OK

	// { SetNX } method example
	resBool, _ := rdb.SetNX(ctx, "key3", "value3", 0).Result()
	fmt.Println(resBool) // -----> true

	resBool, _ = rdb.SetNX(ctx, "key3", "replaced-value", 0).Result()
	fmt.Println(resBool) // -----> false

	// { Get } method example
	resStr, _ = rdb.Get(ctx, "key2").Result()
	fmt.Println(resStr) // -----> value2

	resStr, _ = rdb.Get(ctx, "key4").Result()
	fmt.Println(resStr) // -----> {}

	// { GetSet } method example
	resStr, _ = rdb.GetSet(ctx, "key1", "replaced-value-2").Result()
	fmt.Println(resStr) // -----> older-value

	resStr, _ = rdb.GetSet(ctx, "key5", "new-value").Result()
	fmt.Println(resStr) // -----> {}

	// { MGet } method example
	resInterface, _ := rdb.MGet(ctx, "key1", "key2", "key3", "key4", "key8").Result()
	fmt.Println(resInterface) // ----> [replaced-value-2 value2 value3 value4 <nil>]

	// { MSet } method example
	resStr, _ = rdb.MSet(ctx, "key6", "value6", "key7", "value7").Result()
	fmt.Println(resStr)

	// { Incr, IncrBy } method example
	resCount, _ := rdb.Set(ctx, "count", 0, 0).Result()
	fmt.Println(resCount) // ----> OK

	countValue, _ := rdb.Incr(ctx, "count").Result()
	fmt.Println(countValue) // -----> 1

	countValue, _ = rdb.Incr(ctx, "count").Result()
	fmt.Println(countValue) // -----> 2

	countValue, _ = rdb.IncrBy(ctx, "count", 10).Result()
	fmt.Println(countValue) // -----> 12

	// {Decr, DecrBy } method example
	countValue, _ = rdb.Decr(ctx, "count").Result()
	fmt.Println(countValue) // ------> 11

	countValue, _ = rdb.Decr(ctx, "count").Result()
	fmt.Println(countValue) // ------> 10

	countValue, _ = rdb.DecrBy(ctx, "count", 5).Result()
	fmt.Println(countValue) // -------> 5

}
