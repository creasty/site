package store

import (
	"github.com/garyburd/redigo/redis"

	"github.com/creasty/site/utils"
)

var Redis *redis.Conn

func initRedis() {
	c, err := redis.Dial("tcp", utils.Config.RedisUrl)
	if err != nil {
		panic(err)
	}
	// defer c.Close()

	Redis = &c
}
