package controllers

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Login() fiber.Handler
	SignUp() fiber.Handler
}

type ChatController interface {
	CreateRoom() fiber.Handler
	ListMessageByRoom() fiber.Handler
}
