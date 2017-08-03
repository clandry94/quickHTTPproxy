package proxy

import (
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/ivahaev/go-logger"
	"net"
)

const MaxConnections = 50

type Handler struct {
	NewConnections chan net.Conn
	WorkerCount    int
	Port           string
}

func New(s *spec.ProxySpec) *Handler {
	logger.Info("Creating new proxy handler")

	return &Handler{
		NewConnections: make(chan net.Conn, MaxConnections),
		WorkerCount:    s.HandlerSpec.WorkerCount,
		Port:           s.HandlerSpec.Port,
	}
}

func (h *Handler) Listen() error {
	ln, err := net.Listen("tcp", h.Port)

	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Listening...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			return err
		}

		h.NewConnections <- conn
		logger.Info("Pushed conn onto channel")
	}
}
