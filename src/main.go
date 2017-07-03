package main

import (
	"fmt"

	"github.com/clandry94/quickHTTPproxy/src/handler"
	"github.com/clandry94/quickHTTPproxy/src/queue"
)

func main() {
	connectionHandler := &handler.Handler{
		Queue:       make(queue.Queue, 512),
		WorkerCount: 3,
	}
	fmt.Println(connectionHandler.WorkerCount)
	connectionHandler.HandleConnection()
	fmt.Println("test")
}
