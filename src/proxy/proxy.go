package proxy

import (
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/queue"
	//"net"
)

type Handler struct {
	rankedQueueMap queue.RankedQueueMap
	WorkerCount    int
	Port           int
}

type HandlerSpec struct {
	WorkerCount  int
	QueueConfigs []QueueConfig
	Port         int
}

type QueueConfig struct {
	Tag      string
	Priority int
}

func New(spec *HandlerSpec) *Handler {
	rankedQueueMap := queue.NewRankedQueueMap()

	for _, queueConfig := range spec.QueueConfigs {
		rankedQueue := queue.NewRankedQueue(queueConfig.Tag, queueConfig.Priority)
		rankedQueueMap.Insert(&rankedQueue)
	}

	return &Handler{
		WorkerCount: spec.WorkerCount,
		Port:        spec.Port,
	}
}

func (h *Handler) HandleConnection() {
	fmt.Println("handling")
	//TODO add to the queue, need threadsafe queue methods
}
