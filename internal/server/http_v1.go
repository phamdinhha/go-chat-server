package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/controllers"
	"github.com/phamdinhha/go-chat-server/internal/repositories"
	"github.com/phamdinhha/go-chat-server/internal/services"
)

var RegisterAuthRoutes = func(router *gin.RouterGroup, cfg *config.Config, db *sqlx.DB) {

	userRepo := repositories.NewUserRepo(db)
	authService := services.NewAuthService(userRepo, cfg)
	authController := controllers.NewAuthController(authService)

	router.POST("/login", authController.Login)
	router.POST("/signup", authController.SignUp)
}

var RegisterChatRoutes = func(router *gin.RouterGroup, cfg *config.Config, db *sqlx.DB) {
	chatRepo := repositories.NewChatRepo(db)
	roomRepo := repositories.NewChatRoomRepo(db)
	chatService := services.NewChatService(chatRepo, roomRepo, cfg)
	chatController := controllers.NewChatController(chatService)

	router.POST("/rooms", chatController.CreateRoom)
	router.GET("/rooms/:room_id/messages", chatController.ListMessageByRoom)
}
