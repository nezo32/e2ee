package websocket

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type serverImpl struct {
	address string
	endpoint string
	upgrader websocket.Upgrader
	clients map[uuid.UUID]*websocket.Conn
}

type Server interface {
	Start()
}

type ServerParams struct {
	Address string
	Endpoint string
	Upgrader websocket.Upgrader
}

func NewServer(params *ServerParams) Server {
	return &serverImpl{
		address: params.Address,
		upgrader: params.Upgrader,
		endpoint: params.Endpoint,
		clients: make(map[uuid.UUID]*websocket.Conn),
	}
}

func (s *serverImpl) Start() {
	http.HandleFunc(s.endpoint, s.handleConnection)
	http.ListenAndServe(s.address, nil)
}

func (s *serverImpl) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Upgrade error", "error", err)
		return
	}
	defer conn.Close()

	id := uuid.New()
	s.clients[id] = conn

	log.Info("Client connected", "user", id)

	for {
		mt, _, err := conn.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			log.Info("Client disconnected", "user", id.String())
			break
		}

		log.Info("Received message", "user", id.String())
	}
}