package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yangyulong/secproxy/src/basecomponent/mydatabase"
	"github.com/yangyulong/secproxy/src/basecomponent/myredis"
	"github.com/yangyulong/secproxy/src/controller/secondinfo"
	"github.com/yangyulong/secproxy/src/controller/secondkill"
)

func IninServer() error {
	// config.Parse() //解析配置文件

	//redis
	if err := myredis.InitRedis(); err != nil {
		return fmt.Errorf("init redis err:%s", err)
	}

	//mysql
	if err := mydatabase.InitMysql(); err != nil {
		return fmt.Errorf("init mysql err:%s", err)
	}

	return nil
}

func StartServer() *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/", &secondkill.SecondKill{})
	mux.Handle("/info", &secondinfo.SecondInfo{})
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return &s
}
func main() {
	// pool := &redis.Pool{
	// 	MaxIdle:   4,
	// 	MaxActive: 4,
	// 	Dial: func() (redis.Conn, error) {
	// 		rc, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("yangyulong"))
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return rc, nil
	// 	},
	// 	IdleTimeout: time.Second,
	// }
	// con := pool.Get()
	// str, err := redis.String(con.Do("get", "test"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer con.Close()
	// log.Println("value: ", str)
	err := IninServer()
	if err != nil {
		panic(err)
	}
	s := StartServer()

	//优雅退出
	var closeChannel = make(chan struct{})
	go func() {
		var sign = make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGINT, syscall.SIGKILL)
		<-sign

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		fmt.Println("server will be closed!!!")
		if err := s.Shutdown(ctx); err != nil {
			panic(err)
		}
		close(closeChannel)
	}()
	fmt.Println("server is starting...")
	if err := s.ListenAndServe(); err != nil || err != http.ErrServerClosed {
		panic(err)
	}
	<-closeChannel
	fmt.Println("server had close")
}
