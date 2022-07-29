package redisdb

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

type ConfigRedis struct {
	Host     string
	Port     string
	Database int
	Password string
}

var RedisClient *redis.Client

func ConnectionClient(config ConfigRedis) *redis.Client {
	dsn := fmt.Sprintf("%s:%s", config.Host, config.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: config.Password, // no password set
		DB:       config.Database, // use default DB
	})
	RedisClient = rdb
	return rdb
}
