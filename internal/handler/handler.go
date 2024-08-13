package handler

import (
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	Health(c *fiber.Ctx) error
}

type handler struct {
	service service.IService
}

func BuildHandler(service service.IService) IHandler {
	return &handler{
		service: service,
	}
}

func (h *handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}
