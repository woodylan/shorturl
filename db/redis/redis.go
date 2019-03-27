package redis

import "github.com/go-redis/redis"

var RedisConnect *redis.Client

func Connect() (redisConnect *redis.Client, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = client.Ping().Result()

	redisConnect = client
	RedisConnect = client
	return
}
