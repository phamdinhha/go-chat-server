package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/server"
	"github.com/phamdinhha/go-chat-server/pkg/boot"
	"github.com/phamdinhha/go-chat-server/pkg/postgres"
)

func main() {
	var ctx = context.Background()
	boot.BootstrapDaemons(ctx, StartServer)
}

func StartServer(ctx context.Context) (shutdown boot.Daemon, err error) {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		panic(err)
	}
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	server.RegisterAuthRoutes(ginEngine.Group("/auth"), cfg, db)
	server.RegisterChatRoutes(ginEngine.Group("/chat"), cfg, db)
	server.RegisterWebsocketRoute(ginEngine.Group("/ws"), cfg)

	srvAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	httpServer := &http.Server{
		Addr:    srvAddr,
		Handler: ginEngine,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			panic(err)
		}
	}()

	shutdown = func() {
		<-ctx.Done()

		if err := httpServer.Shutdown(context.Background()); err != nil {
			// s.log.Errorf("Shutdown server error: %v", err)
			fmt.Println(err)
			return
		}

		// s.log.Info("Gracefully shutdown server")
		fmt.Println("Gracefully shutdown server")
	}
	return shutdown, err
}
