package service

import (
	"context"
	"github.com/tuananh3561/go_crm/app/database/db/redisdb"
	"time"
)

func RedisSet(key string, value interface{}, expiration time.Duration) error {
	err := redisdb.RedisClient.Set(context.TODO(), key, value, expiration).Err()
	return err
}

func RedisGet(key string) (string, error) {
	val, err := redisdb.RedisClient.Get(context.TODO(), key).Result()
	return val, err
}

func RedisDel(key string) error {
	err := redisdb.RedisClient.Del(context.TODO(), key).Err()
	return err
}

func RedisHSet(key string, value ...interface{}) error {
	err := redisdb.RedisClient.HSet(context.TODO(), key, value).Err()
	return err
}

func RedisHGet(key string, field string) (string, error) {
	val, err := redisdb.RedisClient.HGet(context.TODO(), key, field).Result()
	return val, err
}

func RedisHDel(key string, fields string) error {
	err := redisdb.RedisClient.HDel(context.TODO(), key, fields).Err()
	return err
}
