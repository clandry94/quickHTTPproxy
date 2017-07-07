package proxy

import (
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/queue"
	"github.com/golang/glog"
	//"net"
)

type HandlerSpec struct {
	WorkerCount  int
	Port         int
	QueueConfigs []QueueConfig
}

type QueueConfig struct {
	Tag      string
	Priority int
}

type Handler struct {
	rankedQueueMap *queue.RankedQueueMap
	WorkerCount    int
	Port           int
}

func New(spec *HandlerSpec) *Handler {
	glog.Info("Creating new proxy handler")
	rqm := queue.NewRankedQueueMap()

	for _, queueConfig := range spec.QueueConfigs {
		rankedQueue := queue.NewRankedQueue(queueConfig.Tag, queueConfig.Priority)
		rqm.Insert(&rankedQueue)
	}

	return &Handler{
		rankedQueueMap: rqm,
		WorkerCount:    spec.WorkerCount,
		Port:           spec.Port,
	}
}

func (h *Handler) HandleConnection() {
	fmt.Println("handling")
	//TODO add to the queue, need threadsafe queue methods
}
