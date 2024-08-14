package config

type Config struct {
	Server  ServerConfig
	Service ServiceConfig
	DB      DBConfig
	Log     LogConfig
}

type ServerConfig struct {
	Port      string `envconfig:"SERVER_PORT"`
	JWTSecret string `envconfig:"JWT_SECRET"`
}

type ServiceConfig struct {
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Database string `envconfig:"DB_DATABASE"`
}

type LogConfig struct {
}
