package handler

import (
	resource "github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"github.com/barisaskaleli/lightweight-bank/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type IHandler interface {
	Health(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Transfer(c *fiber.Ctx) error
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

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Register a new user with the provided details
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		request.RegisterRequest	true	"Register Request"
//	@Success		201				{object}	response.RegisterResponse
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		500				{object}	map[string]interface{}
//	@Router			/register [post]
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

	return c.Status(fiber.StatusCreated).JSON(res)
}

// Login godoc
//
//	@Summary		Login a user
//	@Description	Login a user with the provided credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		request.LoginRequest	true	"Login Request"
//	@Success		200				{object}	response.LoginResponse
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		500				{object}	map[string]interface{}
//	@Router			/login [post]
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

// Transfer godoc
//
//	@Summary		Transfer money between accounts
//	@Description	Transfer money from one account to another
//	@Tags			transaction
//	@Accept			json
//	@Produce		json
//	@Param			transferRequest	body		request.TransferRequest	true	"Transfer Request"
//	@Success		200				{object}	response.TransferResponse
//	@Failure		400				{object}	map[string]interface{}
//	@Failure		401				{object}	map[string]interface{}
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Router			/transfer [post]
func (h *handler) Transfer(c *fiber.Ctx) error {
	var req resource.TransferRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := h.validator.Struct(req)
	if err != nil {
		errors := util.FormatValidationErrors(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accountNumber := claims["account_number"].(string)

	if accountNumber != req.Sender || req.Sender == req.Receiver {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not authorized to transfer from this account.",
		})
	}

	res, err := h.service.Transfer(req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
