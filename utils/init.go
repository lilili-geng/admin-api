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

	go func() {
		log.Printf("%s service running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)

	// SIGINT 用户发送intr 字符(ctrl+c)触发
	// SIGTERM 结束程序
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting Down project %s", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln("%s Shutodwn, causer by :", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout")
	}
	log.Println("%s stop success", srvName)
}
