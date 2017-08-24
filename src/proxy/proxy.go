package proxy

import (
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/clandry94/quickHTTPproxy/src/worker"
	"github.com/ivahaev/go-logger"
	"net"
	"net/http"
)

const MaxConnections = 1000
const MaxRequests = 1000
const MaxWorkers = 20

type Handler struct {
	Connections chan net.Conn
	Requests    chan *http.Request
	WorkerCount int
	Port        string
	WorkerPool  [MaxWorkers]worker.Worker
}

func New(s *spec.ProxySpec) *Handler {
	logger.Info("Creating proxy handler")
	var pool [MaxWorkers]worker.Worker

	conns := make(chan net.Conn, MaxConnections)
	for i := 0; i < MaxWorkers; i++ {
		pool[i] = worker.NewWorker(conns)
		go pool[i].Run()
	}

	return &Handler{
		Connections: conns,
		WorkerCount: s.HandlerSpec.WorkerCount,
		Port:        s.HandlerSpec.Port,
		WorkerPool:  pool,
	}
}

func (h *Handler) Listen() error {
	ln, err := net.Listen("tcp", h.Port)

	if err != nil {
		logger.Error(err)
		return err
	}
	processed := 0
	logger.Info("Listening...")
	logger.Info("Port ", h.Port)
	for {
		logger.Info("Number of connections", len(h.Connections))
		if len(h.Connections) > MaxConnections*0.90 {
			logger.Warn("Approaching maximum connections", len(h.Connections))
		}

		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			return err
		}

		processed++
		logger.Info(processed)
		h.Connections <- conn
		logger.Info("Pushed conn onto channel")
	}
}
