package myredis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	pool *redis.Pool
}

var (
	RedisConf = &RedisClient{}
)

func newPool(server string) *redis.Pool {
	RedisPool := &redis.Pool{
		IdleTimeout: time.Second,
		MaxActive:   10,
		MaxIdle:     5,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("AUTH", "yangyulong")
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return RedisPool
}
func InitRedis() error {
	RedisConf.pool = newPool("127.0.0.1:6379")
	return nil
}

func (rc *RedisClient) do(command string, arges ...interface{}) (interface{}, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	return conn.Do(command, arges...)
}

func (rc *RedisClient) Get(key string) (string, error) {
	return redis.String(rc.do("get", key))
}

func (rc *RedisClient) Set(key string, value string) (string, error) {
	return redis.String(rc.do("SET", key, value))
}
