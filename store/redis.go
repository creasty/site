package store

import (
	"errors"
	"time"

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
	_, err := self.Do("SET", key, value)
	return err
}

func (self *RedisClient) SetWithExpiration(key string, value interface{}, expiration time.Duration) error {
	expiration /= time.Second
	if expiration == 0 {
		return errors.New("expiration is required to be more than 0 seconds")
	}

	self.Send("MULTI")
	self.Send("SET", key, value)
	self.Send("EXPIRE", key, expiration)
	_, err := self.Do("EXEC")
	return err
}

func initRedis() {
	c, err := redis.DialURL(utils.Config.RedisUrl)
	if err != nil {
		panic(err)
	}

	Redis = &RedisClient{c}
}
