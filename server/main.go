package main

import (
	"github.com/nezo32/e2ee/config"
	ws "github.com/nezo32/e2ee/websocket"
	"golang.org/x/net/websocket"
)

func main() {
	cfg := config.InitConfig()

	cfg.LogConfig.SetupLogger()
	cfg.DatabaseConfig.SetupDatabase(config.PostgresDB{})
	
	ws.NewServer(&ws.ServerParams{Address: cfg.ServerConfig.Address, Endpoint: cfg.ServerConfig.Endpoint, Config: &websocket.Config{
		Origin: nil,
	}}).Start()
}