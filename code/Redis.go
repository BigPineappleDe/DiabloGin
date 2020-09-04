package code

import (
	"diabloGin/config"
	"github.com/go-redis/redis"
)

/**
redis核心操作
*/
func RedisInit() *redis.Client{
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.REDISHOST,
		DB:   0, //redis默认有0-15共16个数据库，这里设置操作索引为0的数据库
	})
	_,err:=redisClient.Ping().Result()

	if err!=nil{
		panic(err)
	}
	return redisClient
}
