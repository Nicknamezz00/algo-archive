package main

import (
	"algo-archive/internal"
	"algo-archive/internal/conf"
	"algo-archive/internal/routers"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.Println("Initializing...")
	conf.Initialize()
	log.Println("conf successfully initialized")
	internal.Initialize()
	log.Println("internal successfully initialized")
}

func main() {
	gin.SetMode(conf.ServerSetting.RunMode)

	router := routers.NewRouter()

	srv := &http.Server{
		Addr:           conf.ServerSetting.HTTPIp + ":" + conf.ServerSetting.HTTPPort,
		Handler:        router,
		ReadTimeout:    conf.ServerSetting.ReadTimeOut,
		WriteTimeout:   conf.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	// Graceful server shutdown - https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Algo archive listening on port %v\n", srv.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Blocks until a SIGINT or SIGTERM is passed into the quit channel
	<-quit

	// Shutdown server
	log.Println("Shutdown server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %v\n", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
