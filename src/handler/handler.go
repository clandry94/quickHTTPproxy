package handler

import (
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/queue"
)

type Handler struct {
	Queue       queue.Queue
	WorkerCount int
	// TODO add thread pooling
}

func (h *Handler) HandleConnection() {
	fmt.Println("handling")
	//TODO add to the queue, need threadsafe queue methods
}
