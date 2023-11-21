package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/services"
)

type chatController struct {
	chatService services.ChatService
}

func NewChatController(chatService services.ChatService) ChatController {
	return &chatController{
		chatService: chatService,
	}
}

func (a *chatController) CreateRoom(c *gin.Context) {
	req := dto.RoomReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdRoom, err := a.chatService.CreateRoom(c, req.Name)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdRoom)
	return
}

func (a *chatController) ListMessageByRoom(c *gin.Context) {
	roomID := c.Param("room_id")
	messages, err := a.chatService.ListByRoomId(c, roomID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
	return
}
