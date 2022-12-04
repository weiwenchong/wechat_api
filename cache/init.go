package cache

import (
	"github.com/go-redis/redis"
	"wechat_api/log"
)

var Client *redis.Client

func init() {
	log.Infof("cache init")
	Client = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "redis-77",
		DB:       0,
	})
	log.Infof("cache init succeed")
}
