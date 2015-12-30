package store

import (
	"github.com/garyburd/redigo/redis"

	"github.com/creasty/site/utils"
)

type RedisClient struct {
	redis.Conn
}

var Redis *RedisClient

func (self *RedisClient) Get(key string) (interface{}, error) {
	return self.Do("GET", key)
}

func (self *RedisClient) Set(key string, value interface{}) error {
	return self.Send("SET", key, value)
}

func initRedis() {
	c, err := redis.Dial("tcp", utils.Config.RedisUrl)
	if err != nil {
		panic(err)
	}

	Redis = &RedisClient{c}
}
