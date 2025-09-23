package config

import "github.com/joho/godotenv"

type Config struct {
	DatabaseConfig *dbConfig
	LogConfig      *logConfig
	ServerConfig   *serverConfig
}

func InitConfig() *Config {
	godotenv.Load(".env")

	return &Config{
		DatabaseConfig: initDbConfig(),
		LogConfig:      initLogConfig(),
		ServerConfig:   initServerConfig(),
	}
}
