package main

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"github.com/barisaskaleli/lightweight-bank/internal/handler"
	repo "github.com/barisaskaleli/lightweight-bank/internal/repository"
	resource "github.com/barisaskaleli/lightweight-bank/internal/resource/model"
	"github.com/barisaskaleli/lightweight-bank/internal/router"
	"github.com/barisaskaleli/lightweight-bank/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	_ "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.BuildConfig()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	mysqlInstance, err := config.ConnectMysql(config.MysqlConfig{
		Host:     cfg.DB().Host,
		Port:     cfg.DB().Port,
		User:     cfg.DB().User,
		Password: cfg.DB().Password,
		Database: cfg.DB().Database,
	})

	if err != nil {
		sugar.Errorf("Error while connecting to mysql: %v", err)
		return
	}

	if !mysqlInstance.Database().Migrator().HasTable(&resource.User{}) {
		err = mysqlInstance.Database().AutoMigrate(&resource.User{}, &resource.Transaction{})
		if err != nil {
			sugar.Errorf("Error while migrating: %v", err)
			return
		}
	}

	validator := validator.New()

	fiberConfig := fiber.Config{
		AppName:     "[LightWeight Bank API]",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	app := fiber.New(fiberConfig)

	createRoutes(cfg, sugar, mysqlInstance, validator).RegisterRoutes(app)

	c := make(chan os.Signal, 1)

	go func() {
		if err := app.Listen(cfg.Server().Port); err != nil {
			sugar.Errorf("Error: Bank Application Starting %s", err.Error())
			c <- syscall.SIGTERM
		}
	}()

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Info("Info: Bank Service Gracefully Shutting")
	if gShoutDown := app.Shutdown(); gShoutDown != nil {
		zap.S().Errorf("Error: %v", gShoutDown)
	}
}

func createRoutes(config config.IConfig, logger *zap.SugaredLogger, db config.IMysqlInstance, validator *validator.Validate) router.IRouter {
	repository := repo.BuildRepository(db, config, logger)
	service := service.BuildService(config, logger, repository)
	handler := handler.BuildHandler(service, validator)

	router := router.BuildRouter(config, logger, handler)
	return router
}
