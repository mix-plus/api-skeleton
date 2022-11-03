package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mix-plus/api-skeleton/internal/config"
	"github.com/mix-plus/api-skeleton/internal/handler"
	"github.com/mix-plus/api-skeleton/internal/svc"
	"github.com/mix-plus/core/conf"
	"github.com/mix-plus/core/di"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	if err := conf.MustLoad(*configFile, &c); err != nil {
		panic(err)
	}

	logger := di.Zap()
	server := di.Server()

	svc.Context = svc.NewServiceContext(c)

	gin.SetMode(c.Mode)

	router := gin.New()
	if c.Mode != gin.ReleaseMode {
		handlerFunc := gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s|%s|%d|%s\n",
					params.Method,
					params.Path,
					params.StatusCode,
					params.ClientIP,
				)
			},
			Output: &di.ZapOutput{Logger: logger},
		})
		router.Use(handlerFunc)
	}

	server.Addr = c.Addr
	server.Handler = router

	handler.Load(router)

	// signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("Server shutdown")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			logger.Errorf("Server shutdown error: %s", err)
		}
	}()

	logger.Infof("Server start at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !strings.Contains(err.Error(), "http: Server closed") {
		panic(err)
	}
}
