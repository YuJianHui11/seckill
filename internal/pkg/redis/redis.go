package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"seckill/config"
)

func InitRedis(conf *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Password: conf.Password,
			DB:       conf.DB,
	})
	
	return client, nil
} 