package redis

import (
	"context"
	"log"

	goredis "github.com/redis/go-redis/v9"
)

func NewClient(addr string) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{Addr: addr})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis ping: %v", err) // pinging redis to check if it's up and running
	}
	log.Printf("connected to redis at %s", addr)

	return rdb
}