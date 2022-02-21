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
