package cmd

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/handler"
	"github.com/barisaskaleli/lightweight-bank/internal/router"
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"go.uber.org/zap"
)

func CreateRoutes(config config.IConfig, logger *zap.SugaredLogger) router.IRouter {
	service := service.BuildService(config, logger)
	handler := handler.BuildHandler(service)

	router := router.BuildRouter(logger, handler)
	return router
}
