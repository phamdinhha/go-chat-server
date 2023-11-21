package controllers

import (
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	SignUp(c *gin.Context)
}

type ChatController interface {
	CreateRoom(c *gin.Context)
	ListMessageByRoom(c *gin.Context)
}
