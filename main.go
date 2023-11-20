package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/server"
	"github.com/phamdinhha/go-chat-server/pkg/postgres"
)

func main() {
	if err := startServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func startServer() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		return err
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
	return nil

	// shutdown = func() {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()
	// 	if err := httpServer.Shutdown(ctx); err != nil {
	// 		fmt.Fprintf(os.Stderr, "%v\n", err)
	// 		panic(err)
	// 	}
	// }
	// return shut
}
