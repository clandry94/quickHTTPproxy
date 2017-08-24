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

type Proxy struct {
	Connections chan net.Conn
	Requests    chan *http.Request
	WorkerCount int
	Port        string
	WorkerPool  [MaxWorkers]worker.Worker
}

func New(s *spec.ProxySpec) *Proxy {
	logger.Info("Creating proxy handler")
	var pool [MaxWorkers]worker.Worker

	conns := make(chan net.Conn, MaxConnections)
	for i := 0; i < MaxWorkers; i++ {
		pool[i] = worker.NewWorker(conns)
		go pool[i].Run()
	}

	return &Proxy{
		Connections: conns,
		WorkerCount: s.HandlerSpec.WorkerCount,
		Port:        s.HandlerSpec.Port,
		WorkerPool:  pool,
	}
}

func (p *Proxy) Listen() error {
	ln, err := net.Listen("tcp", p.Port)

	if err != nil {
		logger.Error(err)
		return err
	}
	processed := 0
	logger.Info("Listening on %v", p.Port)

	for {
		logger.Infof("%v connections", len(p.Connections))
		if len(p.Connections) > MaxConnections*0.90 {
			logger.Warnf("Approaching maximum connections: %v", len(p.Connections))
		}

		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			return err
		}

		processed++
		logger.Info(processed)
		p.Connections <- conn
		logger.Info("Pushed conn onto channel")
	}
}
