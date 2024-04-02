package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//err := rdb.Set(ctx, "name", "Djas", time.Hour).Err()
	//if err != nil {
	//	log.Fatal(err)
	//}

	val, err := rdb.Get(ctx, "name").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(val)
	}

}
