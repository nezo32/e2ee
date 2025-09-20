package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nezo32/e2ee/config"
	ws "github.com/nezo32/e2ee/websocket"
)

func main() {
	cfg := config.InitConfig()

	cfg.LogConfig.SetupLogger()
	cfg.DatabaseConfig.SetupDatabase(config.PostgresDB{})
	
	ws.NewServer(&ws.ServerParams{Address: ":8080", Endpoint: "/", Upgrader: websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}}).Start()
}