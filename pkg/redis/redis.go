package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var (
	Redis *redis.Client
)

func InitRedis() {
	host := viper.GetString("server.redis.host")
	port := viper.GetInt("server.redis.port")
	db := viper.GetInt("server.redis.db")
	addr := fmt.Sprintf("%s:%d", host, port)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       db,
	})
}
