package main

import (
	"context"
	"flag"
	"github.com/cuixiaojun001/linkhome/third_party"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/cuixiaojun001/linkhome/cmd/http/bootstrap"
	"github.com/cuixiaojun001/linkhome/cmd/http/router"
	"github.com/cuixiaojun001/linkhome/common/config"
	"github.com/cuixiaojun001/linkhome/common/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "f", "", "配置文件路径: -f conf/prod.yaml")
	flag.Parse()

	if err := bootstrap.SetUp(configFile); err != nil {
		log.Fatal("bootstrap setup failed ", err)
	}
	defer bootstrap.Destroy()

	// 初始化第三方服务
	if err := third_party.Init(config.GetStringMap("third_party")); err != nil {
		logger.Fatalw("init third party failed", "err:", err)
	}

	ginRoute := gin.New()
	// 允许所有来源，你也可以按需设置
	cfg := cors.DefaultConfig()
	cfg.AllowHeaders = append(cfg.AllowHeaders, "Authorization")
	cfg.AllowHeaders = append(cfg.AllowHeaders, "Access-Control-Allow-Origin")
	cfg.AllowAllOrigins = true

	ginRoute.Use(gin.Recovery()).Use(gin.Logger()).Use(cors.New(cfg))

	router.InitRouter(ginRoute)

	// Listen and Server in 0.0.0.0:8080
	serverAddr := config.GetStringMust("listen.http")
	listenPort := 80
	arr := strings.Split(serverAddr, ":")
	if len(arr) == 2 {
		listenPort, _ = strconv.Atoi(arr[1])
	}

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      ginRoute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
			logger.GetAppLogger().Fatalw("listen error", "err", err)
		}
	}()

	logger.GetAppLogger().Infow("Server start", "addr", serverAddr, "port", listenPort)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")
	logger.GetAppLogger().Infow("Shuting down server...")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	//TODO 10s根据需要改动
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
