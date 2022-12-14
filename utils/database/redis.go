package database

import (
	"context"
	"dot/features/products"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func GetRedis() *redis.Client {
	var redisHost = "localhost:6379"
	var redisPassword = ""

	rdb := newRedisClient(redisHost, redisPassword)

	fmt.Println("redis client initialized")

	return rdb
}

func SetDataRedis(save []products.CoreProduct) {
	rdb := GetRedis()
	key := "Product List"
	data := save
	ttl := time.Duration(86400) * time.Second

	// store data using SET command
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
	log.Println("set operation success")

}
