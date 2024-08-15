package router

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	_ "github.com/barisaskaleli/lightweight-bank/docs"
	"github.com/barisaskaleli/lightweight-bank/internal/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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

func (r *router) protected() fiber.Handler {
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
	r.RegisterSwagger(app)
	r.Register(app)
	r.Login(app)
	r.Transfer(app)
}

func (r *router) HealthCheck(router fiber.Router) {
	router.Get("/health", r.protected(), r.handler.Health)
}

func (r *router) RegisterSwagger(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault)
}

func (r *router) Register(router fiber.Router) {
	router.Post("/register", r.handler.Register)
}

func (r *router) Login(router fiber.Router) {
	router.Post("/login", r.handler.Login)
}

func (r *router) Transfer(router fiber.Router) {
	router.Post("/transfer", r.protected(), r.handler.Transfer)
}
