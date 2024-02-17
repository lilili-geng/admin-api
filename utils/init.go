package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, addr string, srvName string) {


	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: r,
	}

	// 启动 HTTP 服务器
	go func() {
		log.Printf("%s service running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// 监听操作系统的终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server:", srvName)

	// 创建一个上下文对象，并设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭 HTTP 服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown HTTP server: %v", err)
	}

	log.Println("Server stopped:", srvName)
}

