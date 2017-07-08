package proxy

import (
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/ivahaev/go-logger"
	"net"
)

type Handler struct {
	rankedQueueMap *queue.RankedQueueMap
	WorkerCount    int
	Port           string
}

func New(s *spec.ProxySpec) *Handler {
	logger.Info("Creating new proxy handler")
	rqm := queue.NewRankedQueueMap()

	for _, queueConfig := range s.HandlerSpec.QueueConfigs {
		rankedQueue := queue.NewRankedQueue(queueConfig.Tag, queueConfig.Priority)
		rqm.Insert(&rankedQueue)
	}

	return &Handler{
		rankedQueueMap: rqm,
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
		h.rankedQueueMap.RankedQueueMap["conor"].Push(conn)
		logger.Info("Pushed conn onto queue")
		logger.Info("Queue size is now: ", h.rankedQueueMap.RankedQueueMap["conor"].Len())
	}
}
