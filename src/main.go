package main

import (
	"fmt"
	"github.com/clandry94/quickHTTPproxy/src/handler"
)

func main() {
	connectionHandler := handler.NewHandler(0)
	fmt.Println(connectionHandler.WorkerCount)
	connectionHandler.HandleConnection()
	fmt.Println("test")
}
