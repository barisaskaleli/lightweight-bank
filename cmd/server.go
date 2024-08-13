package cmd

import (
	"github.com/barisaskaleli/lightweight-bank/config"
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

	/*mysqlInstance, err := config.ConnectMysql(config.MysqlConfig{
		Host:     cfg.DB().Host,
		Port:     cfg.DB().Port,
		User:     cfg.DB().User,
		Password: cfg.DB().Password,
		Database: cfg.DB().Database,
	})

	if err != nil {
		sugar.Errorf("Error: %v", err)
		return
	}*/

	fiberConfig := fiber.Config{
		AppName:     "[LightWeight Bank API]",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}

	app := fiber.New(fiberConfig)

	// Registering Routes
	CreateRoutes(cfg, sugar).RegisterRoutes(app)

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
