package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phamdinhha/go-chat-server/internal/dto"
	"github.com/phamdinhha/go-chat-server/internal/services"
)

type authController struct {
	authService services.AuthenService
}

func NewAuthController(authService services.AuthenService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (a *authController) Login(c *gin.Context) {
	req := dto.LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := a.authService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (a *authController) SignUp(c *gin.Context) {
	req := dto.SignUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := a.authService.SignUp(c, req)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
