package websocket

import (
	"net/http"

	"github.com/charmbracelet/log"
	"golang.org/x/net/websocket"
)

type serverImpl struct {
	address string
	endpoint string
	config *websocket.Config
}

type Server interface {
	Start()
}

type ServerParams struct {
	Address string
	Endpoint string
	Config *websocket.Config
}

func NewServer(params *ServerParams) Server {
	return &serverImpl{
		address: params.Address,
		endpoint: params.Endpoint,
		config: params.Config,
	}
}

func (s *serverImpl) Start() {
	chatHandler := websocket.Server{
	    Handler: s.handleConnection,
	    Config:  *s.config,
	}

	http.Handle(s.endpoint, chatHandler)
	log.Info("Server started", "address", s.address, "endpoint", s.endpoint)
	log.Fatal(http.ListenAndServe(s.address, nil))
}

func (s *serverImpl) handleConnection(conn *websocket.Conn) {
	log.Info("Client connected")

	for {
		if conn.IsClientConn() {
			log.Info("Client disconnected")
			break
		}

		packet, err := ReadPacket(conn)
		if err != nil {
			log.Error(err)
		}

		err = WritePacket(conn, packet)
		if err != nil {
			log.Error(err)
		}
	}
}