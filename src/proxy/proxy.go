package proxy

import (
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"github.com/clandry94/quickHTTPproxy/src/spec"
	"github.com/ivahaev/go-logger"
	//"net"
)

type Handler struct {
	rankedQueueMap *queue.RankedQueueMap
	WorkerCount    int
	Port           int
}

func New(s *spec.HandlerSpec) *Handler {
	logger.Info("Creating new proxy handler")
	rqm := queue.NewRankedQueueMap()

	for _, queueConfig := range s.QueueConfigs {
		rankedQueue := queue.NewRankedQueue(queueConfig.Tag, queueConfig.Priority)
		rqm.Insert(&rankedQueue)
	}

	return &Handler{
		rankedQueueMap: rqm,
		WorkerCount:    s.WorkerCount,
		Port:           s.Port,
	}
}

func (h *Handler) HandleConnection() {
}
