package config

import "os"

type serverConfig struct {
	Address  string
	Endpoint string
}

func initServerConfig() *serverConfig {
	return &serverConfig{
		Address:  os.Getenv("SERVER_PORT"),
		Endpoint: os.Getenv("SERVER_ENDPOINT"),
	}
}