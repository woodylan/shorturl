package redis

import (
	"github.com/go-redis/redis"
	"github.com/astaxie/beego"
)

var RedisConnect *redis.Client

func Connect() (redisConnect *redis.Client, err error) {
	host := beego.AppConfig.String("redis_host")
	port := beego.AppConfig.String("redis_port")
	password := beego.AppConfig.String("redis_password")
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	_, err = client.Ping().Result()

	redisConnect = client
	RedisConnect = client
	return
}
