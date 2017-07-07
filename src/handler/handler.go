package handler

import (
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/queue_map"
)

type Handler struct {
	rankedQueueMap queue_map.RankedQueueMap
	WorkerCount    int
	// TODO add thread pooling
}

type HandlerSpec struct {
	WorkerCount int
	QueueConfig *[]QueueConfig
}

type QueueConfig struct {
	tag      string
	priority int
}

func NewHandler(spec *HandlerSpec) Handler {
	// TODO build queue map
	// assign spec
}

func (h *Handler) HandleConnection() {
	fmt.Println("handling")
	//TODO add to the queue, need threadsafe queue methods
}
