package handlers

import "github.com/gofiber/fiber/v2"

type IObjectHandler interface {
	GetObject(ctx *fiber.Ctx) error
}

type ObjectHandler struct {
}
