package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/services"
	"github.com/phamdinhha/go-chat-server/pkg/http_response"
)

type chatController struct {
	chatService services.ChatService
}

func NewChatController(chatService services.ChatService) ChatController {
	return &chatController{
		chatService: chatService,
	}
}

func (a *chatController) CreateRoom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		room := dto.RoomReq{}
		if err := c.BodyParser(&room); err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		createdRoom, err := a.chatService.CreateRoom(c.Context(), room.Name)
		if err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		return http_response.CtxResponse(c, fiber.StatusOK, createdRoom, nil)
	}
}

func (a *chatController) ListMessageByRoom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roomID := c.Params("room_id")
		messages, err := a.chatService.ListByRoomId(c.Context(), roomID)
		if err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		return http_response.CtxResponse(c, fiber.StatusOK, messages, nil)
	}
}
