package myredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/yangyulong/secproxy/src/config"
)

type RedisClient struct {
	redis.Conn
	Pool redis.Pool
}

var (
	RedisConf = RedisClient{}
	RedisPool = &redis.Pool{}
)

func newPool(server string) *redis.Pool {
	RedisPool := redis.Pool{
		IdleTimeout: 60,
		MaxActive:   10,
		MaxIdle:     5,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return &RedisPool
}
func InitRedis() error {
	RedisPool = newPool(config.SecKillConf.Redisconf.Addr)
	return nil
}

func (rc *RedisClient) Do(command string, arges ...string) {

}
func (client RedisClient) Get(key string) (string, error) {
	return redis.String(client.Do("GET", key))
}

func (client RedisClient) Set(key string, value string) (string, error) {
	return redis.String(client.Do("SET", key, value))
}
