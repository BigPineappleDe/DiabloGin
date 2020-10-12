package code

import (
	"diabloGin/config"
	"github.com/go-redis/redis"
	"sync"
)


/**
redis核心操作
防止并发耗费资源 加入线程锁
*/
var redisDb *redis.Client
var redisOnce sync.Once
func RedisInit() *redis.Client{
	redisOnce.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr: config.REDISHOST,
			DB:   0, //redis默认有0-15共16个数据库，这里设置操作索引为0的数据库
		})
		_,err:=redisClient.Ping().Result()

		if err!=nil{
			panic(err)
		}

		redisDb=redisClient
	})

	return redisDb
}
