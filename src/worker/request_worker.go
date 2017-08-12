package worker

import (
	"crypto/sha1"
	"github.com/ivahaev/go-logger"
	"hash"
	"net/http"
)

type RequestWorker struct {
	id       hash.Hash
	Requests chan *http.Request
	quit     chan bool
}

func NewRequestWorker(buf chan *http.Request) RequestWorker {
	return RequestWorker{
		id:       sha1.New(),
		Requests: buf,
		quit:     make(chan bool),
	}
}

func (rw RequestWorker) stop() {
	go func() {
		rw.quit <- true
	}()
}

func (rw RequestWorker) Run() {
	func() {
		for {

			select {
			case newRequest := <-rw.Requests:
				logger.Info("New request for request worker")
				logger.Info(newRequest)
			case <-rw.quit:
				logger.Info("Quitting request worker", rw.id)
			}
		}
	}()
}
