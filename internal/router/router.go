package router

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type IRouter interface {
	RegisterRoutes(app *fiber.App)
}

type router struct {
	config  config.IConfig
	logger  *zap.SugaredLogger
	handler handler.IHandler
}

func BuildRouter(config config.IConfig, logger *zap.SugaredLogger, handler handler.IHandler) IRouter {
	return &router{
		config:  config,
		logger:  logger,
		handler: handler,
	}
}

func (r *router) jwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: jwtware.SigningKey{Key: []byte(r.config.Server().JWTSecret)},
	})
}

func (r *router) RegisterRoutes(app *fiber.App) {
	r.HealthCheck(app)
	r.Register(app)
	r.Login(app)
}

func (r *router) HealthCheck(router fiber.Router) {
	router.Get("/health", r.jwtMiddleware(), r.handler.Health)
}

func (r *router) Register(router fiber.Router) {
	router.Post("/register", r.handler.Register)
}

func (r *router) Login(router fiber.Router) {
	router.Post("/login", r.handler.Login)
}
