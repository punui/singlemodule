package rediscache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/punui/singlemodule/configs"
	"strconv"
)

var redisClient *redis.Client

func init() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     configs.GetRedisCacheConfigs().ADDRESS,
		Password: configs.GetRedisCacheConfigs().PASSWORD,
		DB:       0,
	})
}
func GetKey(key string) int {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	if val != "" {
		n, _ := strconv.Atoi(val)
		return n
	} else {
		return 1
	}
}

func SetKey(key string, value int) {

	err := redisClient.Set(key, value, 0)

	if err != nil {
		println(err)
	}
}
