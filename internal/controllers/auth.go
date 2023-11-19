package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/services"
	"github.com/phamdinhha/go-chat-server/pkg/http_response"
)

type authController struct {
	authService services.AuthenService
}

func NewAuthController(authService services.AuthenService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.LoginRequest{}
		if err := c.BodyParser(&req); err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		user, err := a.authService.Login(c.Context(), req)
		if err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		return http_response.CtxResponse(c, fiber.StatusOK, user, nil)
	}
}

func (a *authController) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.SignUpRequest{}
		if err := c.BodyParser(&req); err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		user, err := a.authService.SignUp(c.Context(), req)
		if err != nil {
			return http_response.ErrorCtxResponse(c, err)
		}
		return http_response.CtxResponse(c, fiber.StatusOK, user, nil)
	}
}
