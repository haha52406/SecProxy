package config

import (
	"os"
)

var ConfPath = "./conf/server.conf"

var (
	SecKillConf = &Config{}
)

type Config struct {
	Serverconf ServerConf
	Redisconf  RedisConf
	Mysqlconf  MysqlConf
}
type ServerConf struct {
	ServerName string
}
type RedisConf struct {
	Addr     string
	PassWord string
}
type MysqlConf struct{}

func Parse() {
	ConfBody, err := os.ReadFile(ConfPath)
	if err != nil {
		panic(err)
	}
}
