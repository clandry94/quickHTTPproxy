package proxy

import (
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/clandry94/quickHTTPproxy/src/worker"
	"github.com/ivahaev/go-logger"
	"net"
)

const MaxConnections = 50
const MaxNewConnWorkers = 5

type Handler struct {
	NewConnections          chan net.Conn
	WorkerCount             int
	Port                    string
	NewConnectionWorkerPool [MaxNewConnWorkers]worker.SortingWorker
}

func New(s *spec.ProxySpec) *Handler {
	logger.Info("Creating new proxy handler")
	var pool [MaxNewConnWorkers]worker.SortingWorker

	newConns := make(chan net.Conn, MaxConnections)
	for i := 0; i < MaxNewConnWorkers; i++ {
		pool[i] = worker.NewSortingWorker(newConns)
	}

	return &Handler{
		NewConnections: newConns,
		WorkerCount:    s.HandlerSpec.WorkerCount,
		Port:           s.HandlerSpec.Port,
		NewConnectionWorkerPool: pool,
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
		if len(h.NewConnections) > MaxConnections*0.90 {
			logger.Warn("Approaching maximum connections", len(h.NewConnections))
		}

		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			return err
		}

		h.NewConnections <- conn
		logger.Info("Pushed conn onto channel")
	}
}
