package handler

import (
	resource "github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"github.com/barisaskaleli/lightweight-bank/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	Health(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type handler struct {
	service   service.IService
	validator *validator.Validate
}

func BuildHandler(service service.IService, validator *validator.Validate) IHandler {
	return &handler{
		service:   service,
		validator: validator,
	}
}

func (h *handler) Health(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}

func (h *handler) Register(c *fiber.Ctx) error {
	var req resource.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.validator.Struct(req)

	if err != nil {
		errors := util.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res, err := h.service.Register(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) Login(c *fiber.Ctx) error {
	var req resource.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.validator.Struct(req)
	if err != nil {
		errors := util.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res, err := h.service.Login(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
