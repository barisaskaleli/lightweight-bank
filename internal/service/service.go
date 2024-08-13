package service

import (
	"github.com/barisaskaleli/lightweight-bank/config"
	"go.uber.org/zap"
)

type IService interface {
	Test()
}

type service struct {
	config config.IConfig
	logger *zap.SugaredLogger
	// repo
}

func BuildService(config config.IConfig, logger *zap.SugaredLogger) IService {
	return &service{
		config: config,
		logger: logger,
	}
}

func (s *service) Test() {
	s.logger.Info("Test")
}
