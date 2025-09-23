package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type poolImpl struct {
	ID      uuid.UUID
	clients map[uuid.UUID]*websocket.Conn
}
