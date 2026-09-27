package app

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {

	address := os.Getenv("REDIS_ADDR")

	if address == "" {
		address = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	_, err := client.Ping(
		context.Background(),
	).Result()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to connect to redis: %w",
			err,
		)
	}

	return client, nil
}
