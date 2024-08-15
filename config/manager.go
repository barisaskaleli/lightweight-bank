package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type IConfig interface {
	Server() ServerConfig
	Service() ServiceConfig
	DB() DBConfig
}

var GlobalConfig IConfig

type config struct {
	cfg Config
}

func BuildConfig() IConfig {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	GlobalConfig = &config{cfg: cfg}
	return GlobalConfig
}

func (c *config) Server() ServerConfig {
	return c.cfg.Server
}

func (c *config) Service() ServiceConfig {
	return c.cfg.Service
}

func (c *config) DB() DBConfig {
	return c.cfg.DB
}
