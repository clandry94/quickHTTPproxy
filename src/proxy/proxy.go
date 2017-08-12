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
const MaxNewConnWorkers = 5
const MaxRequestWorkers = 5

type Handler struct {
	NewConnections          chan net.Conn
	Requests                chan *http.Request
	WorkerCount             int
	Port                    string
	NewConnectionWorkerPool [MaxNewConnWorkers]worker.SortingWorker
	RequestWorkerPool       [MaxRequestWorkers]worker.RequestWorker
}

func New(s *spec.ProxySpec) *Handler {
	logger.Info("Creating new proxy handler")
	var pool [MaxNewConnWorkers]worker.SortingWorker

	logger.Info("Creating sorting workers")
	newConns := make(chan net.Conn, MaxConnections)
	for i := 0; i < MaxNewConnWorkers; i++ {
		pool[i] = worker.NewSortingWorker(newConns)
		go pool[i].Run()
	}

	logger.Info("Creating request workers")
	var requestWorkerPool [MaxRequestWorkers]worker.RequestWorker
	requests := make(chan *http.Request, MaxRequests)
	for j := 0; j < MaxRequestWorkers; j++ {
		requestWorkerPool[j] = worker.NewRequestWorker(requests)
		go requestWorkerPool[j].Run()
	}

	return &Handler{
		NewConnections: newConns,
		WorkerCount:    s.HandlerSpec.WorkerCount,
		Port:           s.HandlerSpec.Port,
		NewConnectionWorkerPool: pool,
		RequestWorkerPool:       requestWorkerPool,
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
		logger.Info("Number of connections", len(h.NewConnections))
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
