package worker

import (
	"bufio"
	"crypto/sha1"
	"github.com/ivahaev/go-logger"
	"hash"
	"net"
	"net/http"
	"net/url"
)

type Worker struct {
	id          hash.Hash
	client      *http.Client
	Connections chan net.Conn
	quit        chan bool
}

func NewWorker(buf chan net.Conn) Worker {
	return Worker{
		id:          sha1.New(),
		client:      &http.Client{},
		Connections: buf,
		quit:        make(chan bool),
	}
}

func (sw *Worker) stop() {
	go func() {
		sw.quit <- true
	}()
}

func (sw *Worker) Run() {
	func() {
		for {

			select {
			case conn := <-sw.Connections:
				logger.Info("New work for worker")
				err := sw.Handle(conn)
				if err != nil {
					logger.Error(err)
				}
				//logger.Info(req)
				//work.Close()
				//logger.Info("Connection closed")
			case <-sw.quit:
				logger.Info("Closing worker with id %v", sw.id)
			}

		}
	}()
}

func (sw *Worker) Handle(conn net.Conn) error {
	b := bufio.NewReader(conn)
	req, err := http.ReadRequest(b)
	if err != nil {
		return err
	}
	logger.Info(req)
	url, err := url.Parse(req.RequestURI)
	if err != nil {
		return err
	}
	req.URL = url
	req.RequestURI = ""
	resp, err := sw.client.Do(req)
	if err != nil {
		return err
	}

	logger.Info(resp)

	return err
}
