package router

import (
	"github.com/barisaskaleli/lightweight-bank/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type IRouter interface {
	RegisterRoutes(app *fiber.App)
}

type router struct {
	logger  *zap.SugaredLogger
	handler handler.IHandler
}

func BuildRouter(logger *zap.SugaredLogger, handler handler.IHandler) IRouter {
	return &router{
		logger:  logger,
		handler: handler,
	}
}

func (r *router) RegisterRoutes(app *fiber.App) {
	r.HealthCheck(app)
}

func (r *router) HealthCheck(router fiber.Router) {
	router.Get("/health", r.handler.Health)
}
